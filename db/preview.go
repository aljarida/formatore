package db

import (
	"bytes"
    "database/sql"
    "fmt"
    "text/tabwriter"
)

func printRow(tw *tabwriter.Writer, raws []sql.RawBytes) {
	for i, raw := range raws {
		if i > 0 {
			fmt.Fprint(tw, "\t")
		}
		if raw == nil {
			fmt.Fprint(tw, "NULL")
		} else {
			fmt.Fprint(tw, string(raw))
		}
	}
	fmt.Fprintln(tw)
}

func printCols(tw *tabwriter.Writer, cols []string) {
    for i, col := range cols {
        if i > 0 {
            fmt.Fprint(tw, "\t")
        }
        fmt.Fprint(tw, col)
    }
    fmt.Fprintln(tw)
}

// NOTE: This code is provided by ChatGPT and can likely be cleaned up.
func PreviewLastN(db *sql.DB, table string, n int) (string, error) {
    query := fmt.Sprintf("SELECT * FROM %s ORDER BY id DESC LIMIT %d", table, n)
    rows, err := db.Query(query)
    if err != nil {
        return "", err
    }
    defer rows.Close()

    cols, err := rows.Columns()
    if err != nil {
        return "", err
    }

    var buf bytes.Buffer
    tw := tabwriter.NewWriter(&buf, 0, 4, 2, ' ', 0)

	printCols(tw, cols)

    vals := make([]interface{}, len(cols))
    raws := make([]sql.RawBytes, len(cols))
    for i := range vals {
        vals[i] = &raws[i]
    }

    for rows.Next() {
        if err := rows.Scan(vals...); err != nil {
            return "", err
        }
		printRow(tw, raws)
    }

    if err := rows.Err(); err != nil {
        return "", err
    }

    if err := tw.Flush(); err != nil {
        return "", err
    }
    return buf.String(), nil
}
