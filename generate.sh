#!/usr/bin/env bash
OUTPUT=$1
if [ -z $OUTPUT ];then
  OUTPUT=./must_generated.go
fi

FROM=$2
if [ -z $FROM ];then
  FROM=1
fi

TO=$3
if [ -z $TO ];then
  TO=2
fi

generate() {
  echo package must

  for n in $(seq $FROM $TO);do
    echo
    echo "// Must$n panics if the function result contains an error,"
    echo "// otherwise it returns the remainder of the return values excluding the error type."
    echo -n "func Must$n["
    for i in $(seq $n);do
      if [ $i -ne 1 ];then
        echo -n ,
      fi
      echo -n "T$i any"
    done
    echo -n "]("
    for i in $(seq $n);do
      echo -n "t$i T$i,"
    done
    echo -n "err error) ("
    for i in $(seq $n);do
      if [ $i -ne 1 ];then
        echo -n ,
      fi
      echo -n "T$i"
    done
    echo ") {"
    echo "  if err != nil {"
    echo "    panic(err)"
    echo "  }"
    echo -n "  return "
    for i in $(seq $n);do
      if [ $i -ne 1 ];then
        echo -n ,
      fi
      echo -n "t$i"
    done
    echo
    echo "}"
  done
}

generate > $OUTPUT
