package dblogic

const PKeyFieldName = "id"
const PKeyComp = PKeyFieldName + " INTEGER PRIMARY KEY AUTOINCREMENT, "
const UnixDataTimeFieldName = "unix_datetime"
const UnixDatetimeComp = UnixDataTimeFieldName + " INTEGER NOT NULL, "

const Integer = "INTEGER"
const Real = "REAL"
const Text = "TEXT"
const Blob = "BLOB"
