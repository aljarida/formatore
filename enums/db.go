package enums

// A reserved table name that cannot be removed.
// A syntax error will be throw if you attempt to adjust this table.
const SqliteSequence = "sqlite_sequence"

// Necessary field names used in all tables.
const PKeyFieldName = "id"
const UnixDataTimeFieldName = "unix_datetime"

// Necessary components used in all table creation queries.
const PKeyComp = PKeyFieldName + " INTEGER PRIMARY KEY AUTOINCREMENT, "
const UnixDatetimeComp = UnixDataTimeFieldName + " INTEGER NOT NULL, "
