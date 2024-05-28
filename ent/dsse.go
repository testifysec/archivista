// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/in-toto/archivista/ent/dsse"
	"github.com/in-toto/archivista/ent/statement"
)

// Dsse is the model entity for the Dsse schema.
type Dsse struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// GitoidSha256 holds the value of the "gitoid_sha256" field.
	GitoidSha256 string `json:"gitoid_sha256,omitempty"`
	// PayloadType holds the value of the "payload_type" field.
	PayloadType string `json:"payload_type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DsseQuery when eager-loading is set.
	Edges          DsseEdges `json:"edges"`
	dsse_statement *int
	selectValues   sql.SelectValues
}

// DsseEdges holds the relations/edges for other nodes in the graph.
type DsseEdges struct {
	// Statement holds the value of the statement edge.
	Statement *Statement `json:"statement,omitempty"`
	// Signatures holds the value of the signatures edge.
	Signatures []*Signature `json:"signatures,omitempty"`
	// PayloadDigests holds the value of the payload_digests edge.
	PayloadDigests []*PayloadDigest `json:"payload_digests,omitempty"`
	// Metadata holds the value of the metadata edge.
	Metadata []*Metadata `json:"metadata,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
	// totalCount holds the count of the edges above.
	totalCount [4]map[string]int

	namedSignatures     map[string][]*Signature
	namedPayloadDigests map[string][]*PayloadDigest
	namedMetadata       map[string][]*Metadata
}

// StatementOrErr returns the Statement value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DsseEdges) StatementOrErr() (*Statement, error) {
	if e.Statement != nil {
		return e.Statement, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: statement.Label}
	}
	return nil, &NotLoadedError{edge: "statement"}
}

// SignaturesOrErr returns the Signatures value or an error if the edge
// was not loaded in eager-loading.
func (e DsseEdges) SignaturesOrErr() ([]*Signature, error) {
	if e.loadedTypes[1] {
		return e.Signatures, nil
	}
	return nil, &NotLoadedError{edge: "signatures"}
}

// PayloadDigestsOrErr returns the PayloadDigests value or an error if the edge
// was not loaded in eager-loading.
func (e DsseEdges) PayloadDigestsOrErr() ([]*PayloadDigest, error) {
	if e.loadedTypes[2] {
		return e.PayloadDigests, nil
	}
	return nil, &NotLoadedError{edge: "payload_digests"}
}

// MetadataOrErr returns the Metadata value or an error if the edge
// was not loaded in eager-loading.
func (e DsseEdges) MetadataOrErr() ([]*Metadata, error) {
	if e.loadedTypes[3] {
		return e.Metadata, nil
	}
	return nil, &NotLoadedError{edge: "metadata"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Dsse) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case dsse.FieldID:
			values[i] = new(sql.NullInt64)
		case dsse.FieldGitoidSha256, dsse.FieldPayloadType:
			values[i] = new(sql.NullString)
		case dsse.ForeignKeys[0]: // dsse_statement
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Dsse fields.
func (d *Dsse) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case dsse.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int(value.Int64)
		case dsse.FieldGitoidSha256:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gitoid_sha256", values[i])
			} else if value.Valid {
				d.GitoidSha256 = value.String
			}
		case dsse.FieldPayloadType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field payload_type", values[i])
			} else if value.Valid {
				d.PayloadType = value.String
			}
		case dsse.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field dsse_statement", value)
			} else if value.Valid {
				d.dsse_statement = new(int)
				*d.dsse_statement = int(value.Int64)
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Dsse.
// This includes values selected through modifiers, order, etc.
func (d *Dsse) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// QueryStatement queries the "statement" edge of the Dsse entity.
func (d *Dsse) QueryStatement() *StatementQuery {
	return NewDsseClient(d.config).QueryStatement(d)
}

// QuerySignatures queries the "signatures" edge of the Dsse entity.
func (d *Dsse) QuerySignatures() *SignatureQuery {
	return NewDsseClient(d.config).QuerySignatures(d)
}

// QueryPayloadDigests queries the "payload_digests" edge of the Dsse entity.
func (d *Dsse) QueryPayloadDigests() *PayloadDigestQuery {
	return NewDsseClient(d.config).QueryPayloadDigests(d)
}

// QueryMetadata queries the "metadata" edge of the Dsse entity.
func (d *Dsse) QueryMetadata() *MetadataQuery {
	return NewDsseClient(d.config).QueryMetadata(d)
}

// Update returns a builder for updating this Dsse.
// Note that you need to call Dsse.Unwrap() before calling this method if this Dsse
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Dsse) Update() *DsseUpdateOne {
	return NewDsseClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Dsse entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Dsse) Unwrap() *Dsse {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Dsse is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Dsse) String() string {
	var builder strings.Builder
	builder.WriteString("Dsse(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("gitoid_sha256=")
	builder.WriteString(d.GitoidSha256)
	builder.WriteString(", ")
	builder.WriteString("payload_type=")
	builder.WriteString(d.PayloadType)
	builder.WriteByte(')')
	return builder.String()
}

// NamedSignatures returns the Signatures named value or an error if the edge was not
// loaded in eager-loading with this name.
func (d *Dsse) NamedSignatures(name string) ([]*Signature, error) {
	if d.Edges.namedSignatures == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := d.Edges.namedSignatures[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (d *Dsse) appendNamedSignatures(name string, edges ...*Signature) {
	if d.Edges.namedSignatures == nil {
		d.Edges.namedSignatures = make(map[string][]*Signature)
	}
	if len(edges) == 0 {
		d.Edges.namedSignatures[name] = []*Signature{}
	} else {
		d.Edges.namedSignatures[name] = append(d.Edges.namedSignatures[name], edges...)
	}
}

// NamedPayloadDigests returns the PayloadDigests named value or an error if the edge was not
// loaded in eager-loading with this name.
func (d *Dsse) NamedPayloadDigests(name string) ([]*PayloadDigest, error) {
	if d.Edges.namedPayloadDigests == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := d.Edges.namedPayloadDigests[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (d *Dsse) appendNamedPayloadDigests(name string, edges ...*PayloadDigest) {
	if d.Edges.namedPayloadDigests == nil {
		d.Edges.namedPayloadDigests = make(map[string][]*PayloadDigest)
	}
	if len(edges) == 0 {
		d.Edges.namedPayloadDigests[name] = []*PayloadDigest{}
	} else {
		d.Edges.namedPayloadDigests[name] = append(d.Edges.namedPayloadDigests[name], edges...)
	}
}

// NamedMetadata returns the Metadata named value or an error if the edge was not
// loaded in eager-loading with this name.
func (d *Dsse) NamedMetadata(name string) ([]*Metadata, error) {
	if d.Edges.namedMetadata == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := d.Edges.namedMetadata[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (d *Dsse) appendNamedMetadata(name string, edges ...*Metadata) {
	if d.Edges.namedMetadata == nil {
		d.Edges.namedMetadata = make(map[string][]*Metadata)
	}
	if len(edges) == 0 {
		d.Edges.namedMetadata[name] = []*Metadata{}
	} else {
		d.Edges.namedMetadata[name] = append(d.Edges.namedMetadata[name], edges...)
	}
}

// Dsses is a parsable slice of Dsse.
type Dsses []*Dsse
