import gleam/int
import gleam/list.{Continue, Stop}
import gleam/result
import gleam/string
import simplifile

const example_file = "./inputs/day2example.txt"

const file = "./inputs/day2.txt"

type Direction {
  Increasing
  Decreasing
  Unknown
}

type Safety {
  Safe
  Unsafe
}

pub fn day2_pt1(use_example: Bool) {
  get_input(use_example)
  |> list.map(check_safe_skip_idx(_, -1))
  |> list.map(fn(safety) {
    case safety {
      Safe -> 1
      Unsafe -> 0
    }
  })
  |> int.sum
}

pub fn day2_pt2(use_example: Bool) {
  get_input(use_example)
  |> list.map(fn(report) {
    case check_safe_skip_idx(report, -1) {
      Safe -> Safe
      Unsafe -> check_safe_tolerant(report)
    }
  })
  |> list.map(fn(safety) {
    case safety {
      Safe -> 1
      _ -> 0
    }
  })
  |> int.sum
}

fn check_safe_skip_idx(report: List(Int), skip_idx: Int) {
  let #(safety, _, _, _) =
    report
    |> list.fold_until(#(Safe, 0, 0, Unknown), fn(acc, level) {
      let #(_, cur_idx, prev_level, prev_dir) = acc
      case cur_idx == skip_idx {
        // skip this level
        True -> Continue(#(Safe, cur_idx + 1, prev_level, prev_dir))
        False -> {
          let dir = case prev_dir {
            Increasing | Decreasing -> prev_dir
            Unknown ->
              case prev_level {
                // 0 means no prev
                0 -> Unknown
                _ ->
                  case True {
                    _ if level - prev_level > 0 -> Increasing
                    _ -> Decreasing
                  }
              }
          }
          case dir {
            // at first unskipped value
            Unknown -> Continue(#(Safe, cur_idx + 1, level, dir))
            Increasing ->
              case level - prev_level {
                1 | 2 | 3 -> Continue(#(Safe, cur_idx + 1, level, dir))
                _ -> Stop(#(Unsafe, cur_idx + 1, level, dir))
              }
            Decreasing ->
              case prev_level - level {
                1 | 2 | 3 -> Continue(#(Safe, cur_idx + 1, level, dir))
                _ -> Stop(#(Unsafe, cur_idx + 1, level, dir))
              }
          }
        }
      }
    })
  safety
}

fn check_safe_tolerant(report: List(Int)) {
  let #(_, safety) =
    report
    |> list.fold_until(#(0, Unsafe), fn(acc, _level) {
      let #(idx, _) = acc
      case check_safe_skip_idx(report, idx) {
        Safe -> Stop(#(idx + 1, Safe))
        Unsafe -> Continue(#(idx + 1, Unsafe))
      }
    })
  safety
}

fn get_input(use_example: Bool) {
  case use_example {
    True -> example_file
    False -> file
  }
  |> simplifile.read()
  |> result.unwrap("")
  |> string.split("\n")
  |> list.map(fn(line) {
    line
    |> string.split(" ")
    |> list.map(fn(val) {
      val
      |> int.parse()
      |> result.lazy_unwrap(fn() { panic })
    })
  })
}
