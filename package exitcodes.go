package exitcodes

/*
It is a common problem to use exit code to indicate success
or failure of a program or script.
*/

const (
	SUCCESS                   = 0  /* successful termination */
	ERROR_MSG                 = 64 /* base value for error messages */
	CMDLINE_USAGE_ERROR       = 64 /* command line usage error */
	DATA_FORMAT_ERROR         = 65 /* data format error */
	CANT_OPEN_INPUT           = 66 /* cannot open input */
	ADDRESS_UNKOWN            = 67 /* addressee unknown */
	HOSTNAME_UNKNOWN          = 68 /* host name unknown */
	SERVICE_UNAVAILABLE       = 69 /* service unavailable */
	INTERNAL_SW_ERROR         = 70 /* internal software error */
	SYSTEM_ERROR              = 71 /* system error (e.g., can't fork) */
	CRITICCAL_OS_FILE_MISSING = 72 /* critical OS file missing */
	CANT_CREATE_OUTPUT_FILE   = 73 /* can't create (user) output file */
	IO_ERROR                  = 74 /* input/output error */
	TEMP_ERROR                = 75 /* temp failure; user is invited to retry */
	REMOTE_ERROR_IN_PROTOCOL  = 76 /* remote error in protocol */
	PERM_DENIED               = 77 /* permission denied */
	CONFIG_ERROR              = 78 /* configuration error */
)
