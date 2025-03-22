import gleam/dict
import gleam/int
import gleam/list
import gleam/result
import gleam/string
import simplifile

const example_file = "./inputs/day1example.txt"

const file = "./inputs/day1.txt"

fn get_input(use_example: Bool) {
  case use_example {
    True -> example_file
    False -> file
  }
  |> simplifile.read()
  |> result.unwrap("")
  |> string.split("\n")
  |> list.map(fn(s) {
    let assert Ok(#(l_str, r_str)) = string.split_once(s, "   ")
    let assert Ok(l) = int.parse(l_str)
    let assert Ok(r) = int.parse(r_str)
    #(l, r)
  })
  |> list.unzip()
}

pub fn day1_pt1(use_example: Bool) {
  let #(left, right) = get_input(use_example)

  left
  |> list.sort(int.compare)
  |> list.zip(list.sort(right, int.compare))
  |> list.map(fn(t) {
    let #(l, r) = t
    int.absolute_value(l - r)
  })
  |> int.sum
}

fn calc_occurences(tracker: dict.Dict(Int, Int), values: List(Int)) {
  case values {
    [] -> tracker
    [val, ..rest] ->
      case dict.get(tracker, val) {
        Ok(occurences) -> dict.insert(tracker, val, occurences + 1)
        Error(Nil) -> tracker
      }
      |> calc_occurences(rest)
  }
}

pub fn day1_pt2(use_example: Bool) {
  // calculate similarity score
  // add up each num in left list multiplied 
  // by the number of times it appears in the right list
  let #(left, right) = get_input(use_example)
  let left_occurences =
    left
    |> list.fold(dict.new(), fn(accumulator, value) {
      case dict.get(accumulator, value) {
        Ok(occurences) -> dict.insert(accumulator, value, occurences + 1)
        Error(Nil) -> dict.insert(accumulator, value, 1)
      }
    })

  left
  |> list.map(fn(l) { #(l, 0) })
  |> dict.from_list()
  |> calc_occurences(right)
  |> dict.fold(0, fn(accumulator, key, value) {
    let assert Ok(times) = dict.get(left_occurences, key)
    accumulator + times * key * value
  })
}
