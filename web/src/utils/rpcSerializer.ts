
export class RpcSerializer {
    static serialize(response: any) {
        return {
            total: parseInt(response.total, 10),
        }
    }
}
