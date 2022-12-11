package main

import (
    "fmt"
    "os"
	"strconv"
)

func arg() error {
	#引数の判定
    argc := len(os.Args)
    if argc < 2 {
        return fmt.Errorf("引数が足りません。")
    }

	#引数を配列に格納
    for i := 0; i < argc; i++ {
        argv := os.Args
    }

    return nil
}

func main() {
    if err := arg(); err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	os.Exit(1)	
    }
	printf(".intel_syntax noprefix\n");
  	printf(".globl main\n");
  	printf("main:\n");
  	printf("  mov rax, %d\n", strconv.Atoi(argv[1]));
  	printf("  ret\n");
  	return 0;

}


## C言語
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char **argv) {
  if (argc != 2) {
    fprintf(stderr, "引数の個数が正しくありません\n");
    return 1;
  }

  printf(".intel_syntax noprefix\n");
  printf(".globl main\n");
  printf("main:\n");
  printf("  mov rax, %d\n", atoi(argv[1]));
  printf("  ret\n");
  return 0;
}
