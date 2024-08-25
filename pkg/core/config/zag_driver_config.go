package entities

type DriverConfig interface {
	MaxConcurrency() int
	UsePool() bool
	LogPrint() bool
	LogUpload() bool
	SamplingRate() float64
	MustWriteTime() int64
	SetMaxConcurrency(size int)
	SetUsePool(use bool)
	SetLogPrint(driverLogSwitch bool)
	SetLogUpload(openSamplingLog bool)
	SetSamplingRate(samplingRate float64)
	SetMustWriteTime(mustWriteTime int64)
}

type DriverConfigField struct {
	// max tasks running in the same time
	maxConcurrency int
	// use thread pool
	usePool bool
	// print log switch
	logPrint bool
	// upload log switch
	logUpload bool
	// log upload rate (samplingRate = 1 means 1%)
	samplingRate float64
	// graph runtime more than x ms, must upload log
	mustWriteTime int64
}

func (d *DriverConfigField) LogPrint() bool {
	return d.logPrint
}

func (d *DriverConfigField) SetLogPrint(logPrint bool) {
	d.logPrint = logPrint
}

func (d *DriverConfigField) MaxConcurrency() int {
	return d.maxConcurrency
}

func (d *DriverConfigField) SetMaxConcurrency(maxConcurrency int) {
	d.maxConcurrency = maxConcurrency
}

func (d *DriverConfigField) UsePool() bool {
	return d.usePool
}

func (d *DriverConfigField) SetUsePool(usePool bool) {
	d.usePool = usePool
}

func (d *DriverConfigField) LogUpload() bool {
	return d.logUpload
}

func (d *DriverConfigField) SetLogUpload(logUpload bool) {
	d.logUpload = logUpload
}

func (d *DriverConfigField) SamplingRate() float64 {
	return d.samplingRate
}

func (d *DriverConfigField) SetSamplingRate(samplingRate float64) {
	d.samplingRate = samplingRate
}

func (d *DriverConfigField) MustWriteTime() int64 {
	return d.mustWriteTime
}

func (d *DriverConfigField) SetMustWriteTime(mustWriteTime int64) {
	d.mustWriteTime = mustWriteTime
}

func NewDriverConfig() *DriverConfigField {
	return &DriverConfigField{
		maxConcurrency: 10000,
		usePool:        false,
		logPrint:       false,
		logUpload:      false,
		samplingRate:   0.0,
		mustWriteTime:  5000,
	}
}

var _ DriverConfig = (*DriverConfigField)(nil)
