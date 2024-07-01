#[macro_use]
extern crate rocket;

use redis::Commands;
use rocket::serde::json::Json;
use rocket::serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
struct Data {
    Pais: String,
    Texto: String,
}

#[derive(Serialize, Deserialize)]
struct CountryCount {
    Pais: String,
    Contador: i32,
}

#[post("/set", format = "json", data = "<data>")]
async fn set_data(data: Json<Data>) -> Result<&'static str, &'static str> {
    // Crear cliente de redis
    let client =
        redis::Client::open("redis://redis:6379/").map_err(|_| "Failed to create Redis client")?;

    // Conexion a redis
    let mut con = client
        .get_connection()
        .map_err(|_| "Failed to connect to Redis")?;

    // Insertar en redis
    let _: () = con
        .hincr(&data.Pais, &data.Texto, 1)
        .map_err(|_| "Failed to set data in Redis")?;

    Ok("Data set")
}

#[post("/country_count", format = "json", data = "<country_count>")]
async fn set_country_count(
    country_count: Json<CountryCount>,
) -> Result<&'static str, &'static str> {
    // Crear cliente de redis
    let client =
        redis::Client::open("redis://redis:6379/").map_err(|_| "Failed to create Redis client")?;

    // Conexion a redis
    let mut con = client
        .get_connection()
        .map_err(|_| "Failed to connect to Redis")?;

    // Insertar en redis
    let _: () = con
        .hset(
            "country_counts",
            &country_count.Pais,
            country_count.Contador,
        )
        .map_err(|_| "Failed to set country count in Redis")?;

    Ok("Country count set")
}

#[launch]
fn rocket() -> _ {
    rocket::build().mount("/", routes![set_data, set_country_count])
}
