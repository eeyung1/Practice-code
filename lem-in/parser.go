package main

import (
 "fmt"
 "os"
 "strconv"
 "strings"
)

func ReadFile(filename string) (*ParsedData, error) {
 input, err := os.ReadFile(filename)
 if err != nil { return nil, err }
 return Parse(string(input))
}

// Parse keeps the original input for output and validates rooms before links.
func Parse(input string) (*ParsedData, error) {
 parsed := &ParsedData{Source: input}
 rooms := make(map[string]bool)
 coordinates := make(map[[2]int]bool)
 tunnels := make(map[string]bool)
 antCountSeen, linksSeen := false, false
 startSeen, endSeen := false, false
 pending := ""
 for index, raw := range strings.Split(input,"\n") {
  line := strings.TrimSuffix(raw,"\r")
  fail := func(reason string) (*ParsedData,error) {return nil,fmt.Errorf("line %d: %s",index+1,reason)}
  if line == "" {continue}
  if line == "##start" || line == "##end" {
   if !antCountSeen || linksSeen || pending!="" {return fail("misplaced room command")}
   if line=="##start" {if startSeen {return fail("duplicate start")}; startSeen=true;pending="start"} else {if endSeen {return fail("duplicate end")};endSeen=true;pending="end"}
   continue
  }
  if strings.HasPrefix(line,"#") {continue}
  if !antCountSeen {
   ants,err:=strconv.Atoi(line)
   if err!=nil || ants<=0 {return fail("ant count must be a positive integer")}
   parsed.Ants=ants;antCountSeen=true;continue
  }
  parts:=strings.Fields(line)
  if len(parts)==3 {
   if linksSeen {return fail("room after tunnels")}
   name:=parts[0]
   if strings.HasPrefix(name,"L") || strings.HasPrefix(name,"#") || strings.Contains(name,"-") {return fail("invalid room name")}
   x,errX:=strconv.Atoi(parts[1]);y,errY:=strconv.Atoi(parts[2])
   if errX!=nil || errY!=nil {return fail("invalid coordinates")}
   if rooms[name] {return fail("duplicate room name")}
   point:=[2]int{x,y}
   if coordinates[point] {return fail("duplicate room coordinates")}
   room:=Room{Name:name,X:x,Y:y,IsStart:pending=="start",IsEnd:pending=="end"}
   parsed.Rooms=append(parsed.Rooms,room);rooms[name]=true;coordinates[point]=true;pending=""
   continue
  }
  if pending!="" {return fail("room command must be followed by a room")}
  if strings.TrimSpace(line)!=line || len(parts)!=1 {return fail("invalid tunnel")}
  ends:=strings.Split(line,"-")
  if len(ends)!=2 || !rooms[ends[0]] || !rooms[ends[1]] || ends[0]==ends[1] {return fail("invalid tunnel endpoints")}
  key:=tunnelKey(ends[0],ends[1])
  if tunnels[key] {return fail("duplicate tunnel")}
  tunnels[key]=true;linksSeen=true
  parsed.Tunnels=append(parsed.Tunnels,Tunnel{From:ends[0],To:ends[1]})
 }
 if !antCountSeen || !startSeen || !endSeen || pending!="" {return nil,fmt.Errorf("missing ants, start, or end room")}
 return parsed,nil
}

func tunnelKey(a,b string) string {
 if a>b {a,b=b,a}
 return a+"\x00"+b
}

func endpoints(data *ParsedData) (string,string) {
 start,end:="",""
 for _,room:=range data.Rooms {if room.IsStart {start=room.Name};if room.IsEnd {end=room.Name}}
 return start,end
}
