#!/bin/bash

function pprofExportSvg() {

    DIR="$(dirname $2)"
    mkdir -p "$DIR"

    go tool pprof -svg goroutine.test "$1" > "$2"
    exitIffailed
    open -a "Google Chrome" "$2"
}


function exitIffailed() {
    if [ $? -ne 0 ]; then
      printf "指令失敗, return '$?', 程序終止\n"
      exit 1
    fi

}

function runTest() {    
    go test -v -run="^$1$"
}


function runBenchmark () {

    if [[ -z "$1" ]]; then 
        printf "[Error] empty benchmark function name\n"
        exit 1
    fi


    SOURCE="profiles/$1"

    mkdir -p "$SOURCE"

    go test -v -run=^$ -bench="^$1$"  \
        -cpuprofile="$SOURCE/cpu.out"  \
        -memprofile="$SOURCE/mem.out"  \
        -blockprofile="$SOURCE/block.out"
        # -trace=profiles/trace.out 

    exitIffailed

    pprofExportSvg  "$SOURCE/block.out"     "$SOURCE/block.out.svg"
    pprofExportSvg  "$SOURCE/mem.out"     "$SOURCE/mem.out.svg"
    pprofExportSvg  "$SOURCE/cpu.out"     "$SOURCE/cpu.out.svg"
}

function toLowerCase() {
    echo "$( echo $1 | tr [:upper:] [:lower:] )"
}


printf "Please enter the Test or Benchmark function name:\n"
read FUNC


if [[ "$(toLowerCase "$FUNC" )" == benchmark* ]]; then
    runBenchmark "$FUNC"

elif [[ "$(toLowerCase "$FUNC" )" == test* ]]; then
    runTest "$FUNC"

else
    if [[ -z "$FUNC" ]]; then 
        printf "[Error] empty test function name\n"
        exit 1
    fi
    printf "function name must start with either 'Benchmark' or 'Test'\n"
    exit 1
fi



