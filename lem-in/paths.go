package main

import "sort"

// Split each room into entrance and exit nodes. Intermediate room capacity is
// one; reverse edges let later searches reroute an earlier path choice.
type flowEdge struct {
 to, reverse, capacity, cost, initial int
 tunnel bool
}
type flowNetwork [][]flowEdge

func (network *flowNetwork) add(from,to,capacity,cost int,tunnel bool) {
 forward:=flowEdge{to:to,reverse:len((*network)[to]),capacity:capacity,cost:cost,initial:capacity,tunnel:tunnel}
 backward:=flowEdge{to:from,reverse:len((*network)[from]),cost:-cost}
 (*network)[from]=append((*network)[from],forward)
 (*network)[to]=append((*network)[to],backward)
}

// augment sends one unit along a cheapest residual path. Bellman-Ford handles
// the negative reverse costs without relying on a priority queue.
func (network flowNetwork) augment(source,sink int) bool {
 const infinity=int(^uint(0)>>1)/4
 distance:=make([]int,len(network));parentNode:=make([]int,len(network));parentEdge:=make([]int,len(network))
 for i:=range distance {distance[i]=infinity;parentNode[i]=-1}
 distance[source]=0
 for pass:=0;pass<len(network)-1;pass++ {
  changed:=false
  for from,edges:=range network {
   if distance[from]==infinity {continue}
   for index,edge:=range edges {
    if edge.capacity>0 && distance[from]+edge.cost<distance[edge.to] {
     distance[edge.to]=distance[from]+edge.cost;parentNode[edge.to]=from;parentEdge[edge.to]=index;changed=true
    }
   }
  }
  if !changed {break}
 }
 if parentNode[sink]<0 {return false}
 for to:=sink;to!=source; {
  from,index:=parentNode[to],parentEdge[to]
  reverse:=network[from][index].reverse
  network[from][index].capacity--
  network[to][reverse].capacity++
  to=from
 }
 return true
}

func (network flowNetwork) paths(names []string,ids map[string]int,start,end string) [][]string {
 used:=make(map[int][]int)
 for from,edges:=range network {
  for _,edge:=range edges {
   if edge.tunnel && edge.initial-edge.capacity>0 {used[from/2]=append(used[from/2],edge.to/2)}
  }
 }
 paths:=[][]string{}
 for _,next:=range used[ids[start]] {
  path:=[]string{start};current:=next
  for len(path)<=len(names) {
   path=append(path,names[current])
   if names[current]==end {paths=append(paths,path);break}
   if len(used[current])!=1 {break}
   current=used[current][0]
  }
 }
 sort.Slice(paths,func(i,j int)bool {return len(paths[i])<len(paths[j])})
 return paths
}

// pathSets visits the minimum total length path set for each feasible flow size.
func pathSets(graph *Graph,start,end string,limit int,visit func([][]string)) {
 if graph==nil || start==end {return}
 if _,ok:=graph.Neighbors[start];!ok {return}
 if _,ok:=graph.Neighbors[end];!ok {return}
 names:=make([]string,0,len(graph.Neighbors))
 for name:=range graph.Neighbors {names=append(names,name)}
 sort.Strings(names)
 ids:=make(map[string]int)
 for i,name:=range names {ids[name]=i}
 network:=make(flowNetwork,2*len(names))
 for i,name:=range names {
  capacity:=1
  if name==start || name==end {capacity=len(names)}
  network.add(2*i,2*i+1,capacity,0,false)
  for _,neighbor:=range graph.Neighbors[name] {
   target,ok:=ids[neighbor]
   if !ok || name==end || neighbor==start || name==neighbor {continue}
   network.add(2*i+1,2*target,1,1,true)
  }
 }
 for flow:=0;flow<limit && network.augment(2*ids[start]+1,2*ids[end]);flow++ {visit(network.paths(names,ids,start,end))}
}

// FindPaths retains the lesson API, now using residual flow rather than blocking.
func FindPaths(graph *Graph,start,end string) [][]string {
 var paths [][]string
 if graph==nil {return paths}
 pathSets(graph,start,end,len(graph.Neighbors),func(candidate [][]string){paths=candidate})
 return paths
}

func BestPaths(graph *Graph,start,end string,antCount int) [][]string {
 var best [][]string
 bestTurns:=int(^uint(0)>>1)
 pathSets(graph,start,end,antCount,func(paths [][]string){
  counts:=AllocateAnts(paths,antCount)
  turns:=0
  for i,count:=range counts {
   if count>0 && len(paths[i])+count-2>turns {turns=len(paths[i])+count-2}
  }
  if len(paths)>0 && turns<bestTurns {bestTurns=turns;best=paths}
 })
 return best
}
