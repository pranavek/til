struct Metadata {
    key : String,
    value: String,
}

fn main() {
    let mut metadatas : Vec<Metadata> = Vec::new();
    metadatas.push(Metadata{key: "name".to_string(), value: "pranav".to_string()});
    metadatas.push(Metadata{key: "sex".to_string(), value: "m".to_string()});

    for metadata in metadatas {
        let Metadata { key, value } = metadata;
        println!("{} {}",key, value);
    }
}



