// use std::any::Any;
//
// #[derive(PartialEq, Debug)]
// pub enum Pulse {
//     High, Low
// }
//
// impl From<bool> for Pulse {
//     fn from(v: bool) -> Self {
//         match v {
//             false => Pulse::Low,
//             true => Pulse::High,
//         }
//     }
// }
//
// // struct Module<'a> {
// //     connected: Vec<&'a Module<'a>>,
// // }
// // impl Module {
// //     pub fn connect(&mut self, to_connect: &Module) -> ();
// // }
//
// pub trait Module<'a> {
//     fn connect(&mut self, to_connect: &'a dyn Module) -> ();
//
//     fn handle_pulse(&mut self, input: &Pulse) -> Option<Pulse>;
// }
//
// pub struct FlipFlop<'a> {
//     state: bool,
//     connected: Vec<&'a dyn Module>,
// }
//
// impl FlipFlop<'_> {
//     pub fn new() -> Self {
//         FlipFlop {
//             state: false,
//             connected: vec![],
//         }
//     }
// }
//
// impl Module<'_> for FlipFlop<'_> {
//     fn connect(&'_ mut self, to_connect: &'_ dyn Module) -> () {
//         self.connected.push(to_connect)
//     }
//
//     fn handle_pulse(
//         &mut self,
//         input: &Pulse,
//     ) -> Option<Pulse> {
//         // if matches!(input, Pulse::High) {
//         // if let Pulse::High = input {
//         if *input == Pulse::High  {
//             return None;
//         }
//
//         self.state = !self.state;
//         Some(Pulse::from(self.state))
//     }
// }
//
// #[cfg(test)]
// mod tests {
//     use super::*;
//
//     fn test_handle_pulse() {
//         let mut flipper = FlipFlop::new();
//
//         assert_eq!(flipper.handle_pulse(&Pulse::High), None);
//         assert_eq!(flipper.handle_pulse(&Pulse::Low), Some(Pulse::High));
//         assert_eq!(flipper.handle_pulse(&Pulse::Low), Some(Pulse::Low));
//     }
// }
//
// // struct Conjunction {
// //     connected: Vec<dyn Module>,
// //     state: Vec<Pulse>,
// // }
//
// // impl Conjunction {
// //     pub fn new() -> Conjunction {
// //         Conjunction {
// //             connected: vec![],
// //             state: vec![],
// //         }
// //     }
// //
// //     pub fn connect(
// //         &mut self,
// //         to_connect: &dyn Module,
// //     ) -> () {
// //         self.connected.push(to_connect);
// //         self.state.push(Pulse::Low);
// //     }
// // }
// //
// // impl Module for Conjunction {
// //     fn handle_pulse(&mut self, input: &Pulse) -> Option<Pulse> {
// //         for connected_mod in self.connected {
// //         }
// //
// //         None
// //     }
// // }
