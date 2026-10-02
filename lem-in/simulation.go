package main

import (
 "fmt"
 "os"
 "sort"
)

func Solve(data *ParsedData) ([][]string,error) {
 if data==nil || data.Ants<=0 {return nil,fmt.Errorf("invalid ant count")}
 start,end:=endpoints(data)
 if start=="" || end=="" || start==end {return nil,fmt.Errorf("invalid start or end")}
 paths:=BestPaths(BuildGraph(data),start,end,data.Ants)
 if len(paths)==0 {return nil,fmt.Errorf("no path from start to end")}
 return SimulateTurns(paths,data.Ants)
}

// SimulateTurns moves leading ants first so vacated rooms can be reused in the
// same turn, while enforcing one use of each tunnel and one move per ant.
func SimulateTurns(paths [][]string,antCount int) ([][]string,error) {
 if antCount<=0 || len(paths)==0 {return nil,fmt.Errorf("no ants or usable paths")}
 start,end:="",""
 interiors:=make(map[string]bool)
 directSeen:=false
 for _,path:=range paths {
  if len(path)<2 {return nil,fmt.Errorf("invalid path")}
  if start=="" {start,end=path[0],path[len(path)-1]}
  if path[0]!=start || path[len(path)-1]!=end || start==end {return nil,fmt.Errorf("inconsistent path endpoints")}
  if len(path)==2 {if directSeen {return nil,fmt.Errorf("duplicate direct path")};directSeen=true}
  for _,room:=range path[1:len(path)-1] {
   if room==start || room==end || interiors[room] {return nil,fmt.Errorf("paths share or repeat intermediate rooms")}
   interiors[room]=true
  }
 }
 ants:=AssignAnts(paths,antCount)
 turns:=[][]string{}
 finished:=0
 for finished<antCount {
  order:=make([]int,len(ants))
  occupied:=make(map[string]bool)
  for i,ant:=range ants {order[i]=i;if ant.Position>0 && ant.Position<len(ant.Path)-1 {occupied[ant.Path[ant.Position]]=true}}
  sort.SliceStable(order,func(i,j int)bool{return ants[order[i]].Position>ants[order[j]].Position})
  used:=make(map[string]bool)
  moved:=make(map[int]string)
  for _,index:=range order {
   ant:=&ants[index]
   if ant.Position==len(ant.Path)-1 {continue}
   current,next:=ant.Path[ant.Position],ant.Path[ant.Position+1]
   key:=tunnelKey(current,next)
   if used[key] || (next!=end && occupied[next]) {continue}
   if current!=start {delete(occupied,current)}
   ant.Position++;used[key]=true
   if next==end {finished++} else {occupied[next]=true}
   moved[ant.ID]=next
  }
  if len(moved)==0 {return nil,fmt.Errorf("simulation cannot advance")}
  moves:=make([]string,0,len(moved))
  for id:=1;id<=antCount;id++ {if room,ok:=moved[id];ok {moves=append(moves,fmt.Sprintf("L%d-%s",id,room))}}
  turns=append(turns,moves)
 }
 return turns,nil
}

// Simulate remains available for the original lesson call sites.
func Simulate(paths [][]string,antCount int) {
 turns,err:=SimulateTurns(paths,antCount)
 if err!=nil {fmt.Fprintln(os.Stderr,"ERROR: invalid data format,",err);return}
 for _,moves:=range turns {for i,move:=range moves {if i>0 {fmt.Print(" ")};fmt.Print(move)};fmt.Println()}
}
