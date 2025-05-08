use rocket::State;
use crate::AppConfig;  // 메인 모듈에서 정의한 구조체 사용

#[get("/")]
pub fn list() -> &'static str {
    "User list page"
}

#[get("/<id>")]
pub fn profile(id: usize) -> String {
    format!("User profile page for user {}", id)
}

#[get("/settings")]
pub fn settings(config: &State<AppConfig>) -> String {
    format!("User settings with app key: {}", config.app_key)
}
