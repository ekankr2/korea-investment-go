#[macro_use]
extern crate rocket;
use dotenv::dotenv;
use std::env;

mod routes;

struct AppConfig {
    app_key: String,
    app_secret: String,
}

#[launch]
fn rocket() -> _ {
    dotenv().expect("Failed to load .env file");
    let app_key = env::var("APP_KEY").expect("APP_KEY not set");
    let app_secret = env::var("APP_SECRET").expect("SECRET_KEY not set");

    let config = AppConfig {
        app_key,
        app_secret,
    };

    rocket::build()
        .mount(
            "/",
            routes![
                routes::test_route::list,
                routes::test_route::profile,
                routes::test_route::settings
            ],
        )
        .manage(config)
}
