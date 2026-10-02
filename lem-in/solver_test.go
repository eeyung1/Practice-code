package main

import (
 "bytes"
 "fmt"
 "os"
 "path/filepath"
 "strings"
 "testing"
)

func TestSolveBalancesAndMovesLegally(t *testing.T) {
 input := "6\n##start\nS 0 0\nA 1 0\nB 2 0\nC 3 0\nD 4 0\nF 5 0\nG 6 0\n##end\nE 7 0\nS-A\nA-E\nS-B\nB-C\nC-E\nS-D\nD-F\nF-G\nG-E\n"
 data, err := Parse(input)
 if err != nil { t.Fatal(err) }
 turns, err := Solve(data)
 if err != nil { t.Fatal(err) }
 if len(turns) != 4 { t.Fatalf("got %d turns; want 4", len(turns)) }
 checkMoves(t, data, turns)
}

func TestSolveDirectAndNoRoute(t *testing.T) {
 data, err := Parse("3\n##start\nS 0 0\n##end\nE 1 0\nS-E\n")
 if err != nil { t.Fatal(err) }
 turns, err := Solve(data)
 if err != nil { t.Fatal(err) }
 if len(turns) != 3 { t.Fatalf("direct tunnel: got %d turns; want 3",len(turns)) }
 checkMoves(t,data,turns)
 data.Tunnels = nil
 if _,err := Solve(data); err == nil { t.Fatal("expected no-route error") }
}

func TestParseRejectsInvalidInput(t *testing.T) {
 valid := "3\n##start\nS 0 0\n##end\nE 1 0\nS-E\n"
 cases := []string{"", "0\n", "-2\n", strings.Replace(valid,"##start\n","",1),strings.Replace(valid,"E 1 0","E x 0",1), valid+"garbage\n",valid+"E-S\n",strings.Replace(valid,"S-E","S-X",1),strings.Replace(valid,"E 1 0","E 0 0",1),strings.Replace(valid,"S-E","S-S",1),strings.Replace(valid,"E 1 0","Lbad 1 0",1),"3\n##start\n##end\nS 0 0\n",valid+"X 2 0\n", strings.Replace(valid,"##end","##start",1)}
 for i,input := range cases { if _,err:=Parse(input);err==nil {t.Errorf("case %d accepted invalid input",i)} }
}

func TestParseCommentsCRLFAndNegativeCoordinates(t *testing.T) {
 input := "# comment\r\n2\r\n##start\r\n# comment\r\nS -1 -2\r\n##unknown\r\n##end\r\nE 3 4\r\n\r\nS-E\r\n"
 data,err:=Parse(input)
 if err!=nil {t.Fatal(err)}
 if data.Ants!=2 || data.Rooms[0].X!=-1 || !data.Rooms[0].IsStart {t.Fatalf("bad parse: %+v",data)}
}

func TestRunPrintsInputAndMoves(t *testing.T) {
 input:="2\n##start\nS 0 0\n##end\nE 1 0\nS-E\n"
 path:=filepath.Join(t.TempDir(),"map.txt")
 if err:=os.WriteFile(path,[]byte(input),0600);err!=nil {t.Fatal(err)}
 var output bytes.Buffer
 if err:=Run([]string{path},&output);err!=nil {t.Fatal(err)}
 want:=input+"\nL1-E\nL2-E\n"
 if output.String()!=want {t.Fatalf("output = %q; want %q",output.String(),want)}
 if err:=Run(nil,&output);err==nil {t.Fatal("expected usage error")}
 if err:=Run([]string{path+"missing"},&output);err==nil {t.Fatal("expected file error")}
}

// Independently check every turn against rooms and undirected tunnel capacity.
func checkMoves(t *testing.T,data *ParsedData,turns [][]string) {
 t.Helper()
 start,end := endpoints(data)
 positions:=make([]string,data.Ants+1)
 for i:=1;i<=data.Ants;i++ {positions[i]=start}
 tunnels:=map[string]bool{}
 for _,edge:=range data.Tunnels {tunnels[tunnelKey(edge.From,edge.To)]=true}
 for turn,moves:=range turns {
  moved:=map[int]bool{}; used:=map[string]bool{}
  for _,move:=range moves {
   var id int; var room string
   if _,err:=fmt.Sscanf(move,"L%d-%s",&id,&room);err!=nil {t.Fatal(err)}
   if id<1 || id>data.Ants || moved[id] || positions[id]==end {t.Fatalf("invalid ant move on turn %d: %s",turn+1,move)}
   key:=tunnelKey(positions[id],room)
   if !tunnels[key] || used[key] {t.Fatalf("invalid tunnel on turn %d: %s",turn+1,move)}
   moved[id]=true;used[key]=true;positions[id]=room
  }
  occupied:=map[string]bool{}
  for id:=1;id<=data.Ants;id++ {room:=positions[id];if room!=start && room!=end {if occupied[room] {t.Fatalf("collision turn %d",turn+1)};occupied[room]=true}}
 }
 for id:=1;id<=data.Ants;id++ {if positions[id]!=end {t.Fatalf("ant %d never finished",id)}}
}
