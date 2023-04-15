export enum LogLevel {
    Debug = 1,
    Info = 2,
    Warn = 3,
    Error = 4
}

let loglevel: LogLevel = LogLevel.Info

export const setLogLevel = (level: LogLevel) => {
    loglevel = level
}

export const createLogger = (
	name: string, 
	currentLoglevel: LogLevel = loglevel
) => {
    return {
		setLogLevel: (level: LogLevel) => {
			currentLoglevel = level
		},
        info: (...data: any[]) => {
            if (currentLoglevel > LogLevel.Info) {
                return
            }

            console.log(`[${name}]`, ...data)
        },
        error: (...data: any[]) => {
            if (currentLoglevel > LogLevel.Error) {
                return
            }

            console.error(`[${name}]`, ...data)
        },
        warn: (...data: any[]) => {
            if (currentLoglevel > LogLevel.Warn) {
                return
            }

            console.warn(`[${name}]`, ...data)
        },
        debug: (...data: any[]) => {
            if (currentLoglevel > LogLevel.Debug) {
                return
            }

            console.debug(`[${name}]`, ...data)
        },
        child: (childName: string) => {
            return createLogger(`${name}:${childName}`)
        }
    }
}