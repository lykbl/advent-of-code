use std::collections::HashSet;
use std::fs::File;
use std::io;
use std::io::{BufRead, BufReader};
use std::path::Path;
use std::thread::sleep;
use std::time::Duration;
use graph::Edge;
use crate::colors::{clear_console, Color};
use crate::graph::{Direction, shortest_path};

mod graph;
mod colors;

fn graph_from_file<P>(path: P) -> Result<(Vec<Vec<Edge>>, Vec<Vec<usize>>), io::Error>
where P: AsRef<Path>
{
    let lines = match File::open(path) {
        Ok(file)  => BufReader::new(file).lines(),
        Err(err) => return Err(err),
    };

    let heat_loss_map_2d: Vec<Vec<usize>> = lines
        .flatten()
        .map(|line| line.chars().map(|c| c.to_digit(10).unwrap() as usize).collect())
        .collect();

    let height = heat_loss_map_2d.len();
    let width = heat_loss_map_2d.first().unwrap().len();
    let mut heat_loss_graph: Vec<Vec<Edge>> = vec![];
    for (y, row) in heat_loss_map_2d.iter().enumerate() {
        for x in 0..row.len() {
            heat_loss_graph.push(vec![]);
            if x > 0 {
                heat_loss_graph[y * width + x].push(Edge::new(y * width + x - 1, heat_loss_map_2d[y][x - 1], Direction::LEFT));
            }
            if x < width - 1 {
                heat_loss_graph[y * width + x].push(Edge::new(y * width + x + 1, heat_loss_map_2d[y][x + 1], Direction::RIGHT));
            }
            if y > 0 {
                heat_loss_graph[y * width + x].push(Edge::new((y - 1) * width + x, heat_loss_map_2d[y - 1][x], Direction::DOWN));
            }
            if y < height - 1 {
                heat_loss_graph[y * width + x].push(Edge::new((y + 1) * width + x, heat_loss_map_2d[y + 1][x], Direction::UP));
            }
        }
    }

    Ok((heat_loss_graph, heat_loss_map_2d))
}

fn render_graph(
    graph: &Vec<Vec<usize>>,
    path: &HashSet<usize>,
    current_node: &usize,
) {
    // clear_console();
    return;
    println!("------------------------------------------");
    for (x, row) in graph.iter().enumerate() {
        for (y, heat_loss) in row.iter().enumerate() {
            let cell_index = y * row.len() + x;

            let color = if cell_index == *current_node { Color::Yellow  } else if path.contains(&cell_index) { Color::Green } else { Color::Reset };

            print!("{}{}", color, heat_loss);
            print!("{}", Color::Reset);
        }
        println!();
    }
    sleep(Duration::from_millis(5));
}

fn main() {
    let (heat_loss_graph, heat_loss_map_2d) = graph_from_file(Path::new("input.txt")).unwrap();

    if let Some(shortest) = shortest_path(
        &heat_loss_map_2d,
        &heat_loss_graph,
        0, heat_loss_graph.len() - 1,
        // Some(render_graph),
        render_graph,
    ) {
        println!("Result: {}", shortest);
    }
}
