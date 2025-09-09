// schema.pg.hcl  (atlasgo hcl)

schema "public" {
  comment = "schema comment"
}

table "mailboxes" {
  schema = schema.public
  column "id" {
    null = false
    type = int
    identity {
      generated = ALWAYS
      start = 1
     }
  }
  column "mailbox" {
    null = false
    type = varchar(255)
  }
  column "pwhash" {
    type = varchar(255)
  }
  primary_key {
    columns = [ column.id ]
  }
}
