use std::fs::File;
use std::io;
use std::io::{BufRead, BufReader};
use std::path::Path;
use graph::Edge;
use crate::graph::shortest_path;

mod graph;

fn graph_from_file<P>(path: P) -> Result<Vec<Vec<Edge>>, io::Error>
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
                heat_loss_graph[y * width + x].push(Edge::new(y * width + x - 1, heat_loss_map_2d[y][x - 1]));
            }
            if x < width - 1 {
                heat_loss_graph[y * width + x].push(Edge::new(y * width + x + 1, heat_loss_map_2d[y][x + 1]))
            }
            if y > 0 {
                heat_loss_graph[y * width + x].push(Edge::new((y - 1) * width + x, heat_loss_map_2d[y - 1][x]));
            }
            if y < height - 1 {
                heat_loss_graph[y * width + x].push(Edge::new((y + 1) * width + x, heat_loss_map_2d[y + 1][x]));
            }
        }
    }

    Ok(heat_loss_graph)
}

fn main() {
    let heat_loss_graph = graph_from_file(Path::new("debug.txt")).unwrap();

    if let Some(shortest) = shortest_path(&heat_loss_graph, 0, heat_loss_graph.len() - 1) {
        println!("Result: {}", shortest);
    }
}
