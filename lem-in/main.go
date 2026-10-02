package main

import (
 "fmt"
 "io"
 "os"
 "strings"
)

func Run(args []string,output io.Writer) error {
 if len(args)!=1 {return fmt.Errorf("usage: go run . <map-file>")}
 data,err:=ReadFile(args[0])
 if err!=nil {return fmt.Errorf("ERROR: invalid data format, %w",err)}
 turns,err:=Solve(data)
 if err!=nil {return fmt.Errorf("ERROR: invalid data format, %w",err)}
 // Solve first: invalid maps produce no partial map or movement output.
 if _,err=io.WriteString(output,data.Source);err!=nil {return err}
 if !strings.HasSuffix(data.Source,"\n") {if _,err=io.WriteString(output,"\n");err!=nil {return err}}
 if _,err=io.WriteString(output,"\n");err!=nil {return err}
 for _,moves:=range turns {if _,err=fmt.Fprintln(output,strings.Join(moves," "));err!=nil {return err}}
 return nil
}

func main() {
 if err:=Run(os.Args[1:],os.Stdout);err!=nil {fmt.Fprintln(os.Stderr,err);os.Exit(1)}
}
