mod array;
mod data_type;
mod hashset_collector;
mod index_reader;
mod index_reader_c;
mod index_writer;
mod index_writer_c;
mod linkedlist_collector;
mod util;
mod util_c;
mod vec_collector;

pub fn add(left: usize, right: usize) -> usize {
    left + right
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = add(2, 2);
        assert_eq!(result, 4);
    }
}
