/**
 * 格式化日期时间字符串
 * @param dateString - ISO格式的日期字符串
 * @param format - 输出格式：
 *   - 'datetime' (默认): '2025-01-03 14:25:38'
 *   - 'date': '2025-01-03'
 *   - 'time': '14:25:38'
 *   - 'friendly': '刚刚'/'x分钟前'/'x小时前'/'x天前'等
 * @returns 格式化后的日期字符串
 */
export function formatDate(dateString: string, format: 'datetime' | 'date' | 'time' | 'friendly' = 'datetime'): string {
    const date = new Date(dateString)

    // 处理无效日期
    if (isNaN(date.getTime())) {
        return '无效日期'
    }

    // 补零函数
    const pad = (n: number): string => n.toString().padStart(2, '0')

    // 基础日期部分
    const year = date.getFullYear()
    const month = pad(date.getMonth() + 1)
    const day = pad(date.getDate())
    const hours = pad(date.getHours())
    const minutes = pad(date.getMinutes())
    const seconds = pad(date.getSeconds())

    switch (format) {
        case 'date':
            return `${year}-${month}-${day}`

        case 'time':
            return `${hours}:${minutes}:${seconds}`

        case 'friendly': {
            const now = new Date()
            const diff = now.getTime() - date.getTime()
            const minutes = Math.floor(diff / 60000)
            const hours = Math.floor(minutes / 60)
            const days = Math.floor(hours / 24)

            if (minutes < 1) return '刚刚'
            if (minutes < 60) return `${minutes}分钟前`
            if (hours < 24) return `${hours}小时前`
            if (days < 30) return `${days}天前`
            if (days < 365) return `${Math.floor(days / 30)}个月前`
            return `${Math.floor(days / 365)}年前`
        }

        case 'datetime':
        default:
            return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`
    }
}

/**
 * 使用示例:
 * const isoString = '2025-01-03T06:25:38.568+00:00'
 * 
 * formatDate(isoString)                // '2025-01-03 14:25:38' (假设东八区)
 * formatDate(isoString, 'date')        // '2025-01-03'
 * formatDate(isoString, 'time')        // '14:25:38'
 * formatDate(isoString, 'friendly')    // '刚刚'/'x分钟前'等
 */ 