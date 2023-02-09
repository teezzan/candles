package data

var (
	OpenFieldName   OHLCFieldName = "OPEN"
	HighFieldName   OHLCFieldName = "HIGH"
	LowFieldName    OHLCFieldName = "LOW"
	CloseFieldName  OHLCFieldName = "CLOSE"
	SymbolFieldName OHLCFieldName = "SYMBOL"
	UnixFieldName   OHLCFieldName = "UNIX"

	//DefaultOHLCFieldIndexes defines the default OHLC Field indexes.
	DefaultOHLCFieldIndexes = OHLCFieldIndexes{
		Open: FieldIndex{
			Name: OpenFieldName,
		},
		High: FieldIndex{
			Name: HighFieldName,
		},
		Low: FieldIndex{
			Name: LowFieldName,
		},
		Close: FieldIndex{
			Name: CloseFieldName,
		},
		Symbol: FieldIndex{
			Name: SymbolFieldName,
		},
		Unix: FieldIndex{
			Name: UnixFieldName,
		},
	}
)
