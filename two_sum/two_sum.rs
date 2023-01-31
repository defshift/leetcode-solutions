use std::{collections::HashMap, ops::ControlFlow};

fn two_sum_naive_with_iterators(nums: Vec<i32>, target: i32) -> Vec<i32> {
    let res = nums.iter().enumerate().try_for_each(|(index_first, item)| {
        nums.iter()
            .enumerate()
            .skip(index_first + 1)
            .try_for_each(|(index_second, inner_item)| {
                if item + inner_item == target {
                    ControlFlow::Break(vec![index_first as i32, index_second as i32])
                } else {
                    ControlFlow::Continue(())
                }
            })
    });

    match res {
        ControlFlow::Break(res) => res,
        ControlFlow::Continue(_) => vec![],
    }
}

fn two_sum_hashmap(nums: Vec<i32>, target: i32) -> Vec<i32> {
    let mut hm = HashMap::<i32, i32>::with_capacity(nums.len());

    nums.iter().enumerate().for_each(|(index, value)| {
        hm.insert(*value, index as i32);
    });

    let res = nums.iter().enumerate().try_for_each(|(index, value)| {
        if let Some(inner_index) = hm.get(&(target - value)) {
            if *inner_index as usize == index {
                ControlFlow::Continue(())
            } else {
                ControlFlow::Break(vec![index as i32, *inner_index])
            }
        } else {
            ControlFlow::Continue(())
        }
    });

    match res {
        ControlFlow::Break(res) => res,
        ControlFlow::Continue(_) => vec![],
    }
}

fn two_sum_hashmap_advanced(nums: Vec<i32>, target: i32) -> Vec<i32> {
    let mut hm = HashMap::<i32, i32>::with_capacity(nums.len());

    let res = nums
        .iter()
        .enumerate()
        .try_for_each(|(index, value)| match hm.get(&value) {
            Some(inner_index) => ControlFlow::Break(vec![index as i32, *inner_index as i32]),
            None => {
                hm.insert(target - value, index as i32);
                ControlFlow::Continue(())
            }
        });

    match res {
        ControlFlow::Break(res) => res,
        ControlFlow::Continue(_) => vec![],
    }
}

fn main() {
    // stub main
}
