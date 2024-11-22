mod pan_orbit_camera;

use std::fs;
use std::io::{BufRead, BufReader};
use std::ops::{Add, Sub};
use std::path::Path;
use bevy::color::palettes::basic::BLACK;
use bevy::color::palettes::css::SILVER;
use bevy::prelude::*;
use crate::pan_orbit_camera::{PanOrbitCameraPlugin};

fn main() {
    App::new()
        .add_plugins((
            DefaultPlugins.set(ImagePlugin::default_nearest()),
            PanOrbitCameraPlugin,
            // #[cfg(not(target_arch = "wasm32"))]
            // WireframePlugin,
        ))
        .add_systems(Startup, place_bricks)
        .run();
}

fn place_bricks(
    mut commands: Commands,
    mut meshes: ResMut<Assets<Mesh>>,
    // mut images: ResMut<Assets<Image>>,
    mut materials: ResMut<Assets<StandardMaterial>>,
) {
    let filename = "debug.txt";

    let file = fs::File::open(Path::new(filename));
    if file.is_err() {
        return println!();
    }

    let file = file.expect("Failed to open");
    let lines = BufReader::new(file).lines();

    let default_material = materials.add(StandardMaterial {
        base_color: Color::linear_rgb(0.0, 0.0, 0.0),
        ..default()
    });

    // let (start, end) = (Vec3::new(0.0, 0.0, 0.0), Vec3::new(1.0, 1.0, 1.0));
    // let cube = meshes.add(Cuboid::from_corners(start, end));
    // commands.spawn((
    //     PbrBundle {
    //         mesh: cube,
    //         material: default_material.clone(),
    //         transform: Transform::from_xyz(start.x, 1.0, start.z),
    //         ..default()
    //     },
    // ));

    for line in lines.flatten() {
        if let Some((start, end)) = line.split_once("~").map(|(start, end)| (parse_vec3(start), parse_vec3(end))) {
            let cube = meshes.add(Cuboid::from_corners(start, end.add(1.0)));
            commands.spawn((
                PbrBundle {
                    mesh: cube,
                    material: default_material.clone(),
                    transform: Transform::from_xyz(start.x, start.y, start.z),
                    ..default()
                },
            ));
        }
    }

    commands.spawn(PointLightBundle {
        point_light: PointLight {
            shadows_enabled: true,
            intensity: 10_000_000.,
            range: 100.0,
            shadow_depth_bias: 0.2,
            ..default()
        },
        transform: Transform::from_xyz(8.0, 16.0, 8.0),
        ..default()
    });

    // ground plane
    commands.spawn(PbrBundle {
        mesh: meshes.add(Plane3d::default().mesh().size(50.0, 50.0).subdivisions(10)),
        material: materials.add(Color::from(SILVER)),
        ..default()
    });

    // commands.spawn(Camera3dBundle {
    //     transform: Transform::from_xyz(0.0, 7.,  5.0).looking_at(Vec3::new(0., 1., 0.), Vec3::Y),
    //     ..default()
    // });
}

fn parse_vec3(value: &str) -> Vec3 {
    let values: [f32;3] = value
        .split(",")
        .map(|s| s.parse::<f32>().expect("Failed to parse"))
        .collect::<Vec<f32>>()
        .try_into().expect("Expected 3");

    Vec3 {
        x: values[0],
        y: values[2],
        z: values[1],
    }
}
