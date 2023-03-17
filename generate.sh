#!/usr/bin/env bash
OUTPUT=$1
if [ -z $OUTPUT ];then
  OUTPUT=./must_generated
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

generate_test() {
  echo package must
  echo
  echo 'import "fmt"'

  for n in $(seq $FROM $TO);do
    echo
    echo "func ExampleMust$n() {"
    echo "  defer func() {"
    echo "    if err := recover(); err != nil {"
    echo "      fmt.Println(err)"
    echo "    }"
    echo "  }()"
    echo
    echo -n "  f := func() ("
    for i in $(seq $n);do
      echo -n int,
    done
    echo -n "error) { return "
    for i in $(seq $n);do
      echo -n 0,
    done
    echo "fmt.Errorf(\"Must$n error\") }"
    echo
    echo "  Must$n(f())"
    echo "// Output:"
    echo "// Must$n error"
    echo "}"
  done
}

generate | gofmt > "${OUTPUT}.go"
generate_test | gofmt > "${OUTPUT}_test.go"
