use std::cmp::Ordering;
use std::collections::{BinaryHeap, HashSet};
use std::thread::sleep;
use std::time::Duration;
use itertools::Itertools;

#[derive(Copy, Clone, Eq, PartialEq, Debug, Hash)]
pub enum Direction {
    UP, RIGHT, DOWN, LEFT
}

impl Direction {
    fn is_opposite(&self, direction: Direction) -> bool
    {
        match (self, direction) {
            (Direction::UP, Direction::DOWN) | (Direction::RIGHT, Direction::LEFT) | (Direction::LEFT, Direction::RIGHT) | (Direction::DOWN, Direction::UP) => true,
            _ => false,
        }
    }
}

// #[derive(Copy, Clone, Eq, PartialEq)]
#[derive(Eq, PartialEq)]
struct State {
    cost: usize,
    position: usize,
    steps_taken: usize,
    last_dir: Option<Direction>,
    // path: HashSet<usize>,
    path: Vec<usize>,
}

// The priority queue depends on `Ord`.
// Explicitly implement the trait so the queue becomes a min-heap
// instead of a max-heap.
impl Ord for State {
    fn cmp(&self, other: &Self) -> Ordering {
        // Notice that we flip the ordering on costs.
        // In case of a tie we compare positions - this step is necessary
        // to make implementations of `PartialEq` and `Ord` consistent.
        other.cost.cmp(&self.cost)
            .then_with(|| other.position.cmp(&self.position))
    }
}

// `PartialOrd` needs to be implemented as well.
impl PartialOrd for State {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

// Each node is represented as a `usize`, for a shorter implementation.
#[derive(Debug)]
pub struct Edge {
    node: usize,
    cost: usize,
    direction: Direction,
}

impl Edge {
    pub fn new(node: usize, cost: usize, direction: Direction) -> Self {
        Edge { node, cost, direction }
    }
}

// const MAX_STEPS: usize = 3;
const MAX_STEPS: usize = 10;
const MIN_STEPS: usize = 4;

// Dijkstra's shortest path algorithm.

// Start at `start` and use `dist` to track the current shortest distance
// to each node. This implementation isn't memory-efficient as it may leave duplicate
// nodes in the queue. It also uses `usize::MAX` as a sentinel value,
// for a simpler implementation.

// trait OptionalCb {
//     fn render(&self, graph: &Vec<Vec<usize>>, path: &HashSet<Edge>, node: &usize);
// }
//
// impl OptionalCb for Option<fn(&Vec<Vec<usize>>, &HashSet<usize>, &usize)> {
//     fn render(&self, graph: &Vec<Vec<usize>>, path: &HashSet<Edge>, node: &usize) {
//         if *self.is_some() {
//             self(graph, path, node);
//         }
//     }
// }

pub fn shortest_path (
    original: &Vec<Vec<usize>>,
    adj_list: &Vec<Vec<Edge>>,
    start: usize,
    goal: usize,
    render_callback: fn(&Vec<Vec<usize>>, &Vec<usize>, &usize),
) -> Option<usize>
{
    let mut dist: Vec<usize> = (0..adj_list.len()).map(|_| usize::MAX).collect();
    let mut heap = BinaryHeap::new();
    let mut visited = HashSet::new();
    dist[start] = 0;
    heap.push(State { cost: 0, position: start, steps_taken: 0, last_dir: None, path: Vec::from([start]) });

    while let Some(State { cost, position, steps_taken, last_dir, path }) = heap.pop() {
        if visited.contains(&(position, last_dir, steps_taken)) {
            continue;
        }
        visited.insert((position, last_dir, steps_taken));

        if position == goal && steps_taken >= MIN_STEPS {
            render_callback(original, &path, &position);

            return Some(cost);
        }

        for edge in &adj_list[position] {
            if steps_taken < MIN_STEPS && last_dir.is_some_and(|d| edge.direction != d) {
                continue;
            }
            if steps_taken >= MAX_STEPS && last_dir.is_some_and(|d| edge.direction == d) {
                continue;
            }
            if last_dir.is_some_and(|d| edge.direction.is_opposite(d)) {
                continue;
            }

            let steps_taken = if last_dir.is_some_and(|d| d == edge.direction) { steps_taken } else { 0 };
            let mut new_path = path.clone();
            new_path.push(edge.node);
            let next = State { cost: cost + edge.cost, position: edge.node, last_dir: Some(edge.direction), steps_taken: steps_taken + 1, path: new_path };
            heap.push(next);
        }
    }

    // Goal not reachable
    None
}
