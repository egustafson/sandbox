// atlas.hcl -- configuration

variable "dbuser" {
  type = string
  //default = getenv("ATLAS_DBUSER")
  default = "postgres"
}

variable "dbpass" {
  type = string
  //default = getenv("ATLAS_DBPASS")
  default = "dev"
}

env "local" {
  src = "file://schema.pg.hcl"
  url = "postgres://${var.dbuser}:${var.dbpass}@localhost:5432/postgres?sslmode=disable"
  dev = "docker://postgres/17/dev"
}
