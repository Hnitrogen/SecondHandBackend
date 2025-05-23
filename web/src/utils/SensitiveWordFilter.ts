class SensitiveWordFilter {
    private sensitiveWordMap: Map<string, any>;

    constructor() {
        this.sensitiveWordMap = new Map();
        this.initSensitiveWords();
    }

    // 初始化敏感词库
    private initSensitiveWords() {
        // 这里可以从配置文件或API加载敏感词列表
        // source: https://github.com/adlered/DangerousSpamWords/blob/master/DangerousSpamWords/General_SpamWords_V1.0.1_CN.min.txt
        const sensitiveWords = [
            "18禁", "a片", "caonima", "cnm", "caotama", "cao你", "cao你妈", "fa轮", "hjt", "jb", "jzm", "nmsl", "sb", "vpn",
            "zhengfu", "zheng府", "zf", "zedong", "亂倫", "噴精", "姦淫", "屄", "屄毛", "幹炮", "幹砲", "擠乳汁", "溫家寶",
            "無毛穴", "獸交", "爱女人", "爱液", "扒穴", "拔屄自拍", "白虎阴穴", "白虎少妇", "白浆四溅", "包二奶", "薄熙来",
            "薄码", "暴奸", "爆乳娘", "爆草", "爆操", "被操", "被插", "被干", "逼痒", "逼奸", "博彩", "擦你妈", "操我",
            "操死", "操死你", "操你妈", "操你奶", "操你姐", "操逼", "操你祖宗", "操你大爷", "操你妹", "操穴", "操屄",
            "操他妈", "操你全家", "操烂", "操妻", "操你嘴", "操死她", "操b", "操它妈", "曹刚川", "草你祖宗", "草你大爷",
            "草你妈", "草你吗", "草bi", "草她妈", "草拟吗", "草你娘", "草他妈", "厕所盗摄", "厕奴", "插阴茎", "插逼",
            "插妹妹", "插穴止痒", "插死她", "插比", "插b", "插你", "陈同海", "陈水扁", "陈良宇", "成人片", "成人网站",
            "成人图", "成人电", "成人文", "成人视", "成人自拍", "成人小", "惩公安", "吃鸡巴", "抽插", "臭鸡八", "臭鸡吧",
            "床上写真", "催情藥", "催情药", "催眠水", "催情粉", "大奶子", "大肉棒", "大奶头", "大明运气咒", "大力抽送",
            "大血逼", "大雞巴", "大鸡巴", "大傻b", "戴秉国", "盗撮", "邓小瓶", "邓小平", "邓爷爷", "东北独立", "杜世成",
            "杜德印", "法lun", "法伦功", "法维权", "法一轮", "法正乾", "法车仑", "法轮佛", "法轮", "翻墙", "反共复清",
            "反华示威", "肥逼", "粉穴", "风艳阁", "干你妈", "干的爽", "干穴", "干你娘", "干你妹", "干死你", "干你全家",
            "肛交", "肛门", "肛门拳交", "给你爽", "根达亚文明", "共残主义", "狗娘养", "灌满精液", "郭金龙", "国峰",
            "国锋", "含屌", "喝血社会", "黑毛屄", "胡云松", "胡王八", "胡谨涛", "胡錦濤", "胡海峰", "胡主席", "胡春华",
            "胡温", "虎精逃", "华国", "黄色电影", "回良玉", "混蛋", "激情电", "激情妹", "激情炮", "激情短", "激情小说",
            "鸡巴", "鸡奸", "集体自杀", "挤乳汁", "几吧", "妓女", "家宝", "奸幼", "奸杀", "奸污", "践货", "贱人", "贱比",
            "江泽民", "江澤民", "江x", "江某某", "脚奴", "街头扒衣", "金毛穴", "锦涛", "精子射在", "警察说保", "警方包庇",
            "警车雷达", "警察殴打", "警察的幌", "就去日", "菊花洞", "巨乳", "恐怖份子", "恐怖分子", "抠穴", "口淫",
            "口交", "口活", "口内爆射", "狂乳激揺", "拉登", "浪逼", "浪叫", "雷管", "李小鹏", "李鹏", "李克强", "李洪志",
            "李世民", "習近平", "莲花逼", "炼大法", "梁光烈", "两会又三", "两会代", "刘延东", "露逼", "露b", "乱伦类",
            "乱奸", "乱伦小", "轮子功", "轮奸", "轮功", "伦理毛", "伦理电影", "伦理片", "伦理大", "裸舞视", "裸体", "裸聊网",
            "妈个逼", "妈了个逼", "妈了逼", "麻果配", "麻果丸", "麻古", "玛雅历法", "卖淫", "满狗", "肏屄", "氓培训",
            "猫贼洞", "毛主席", "毛遮洞", "毛则东", "毛泽东", "毛贼东", "毛澤東", "美女高潮", "美艳少妇", "妹按摩", "妹上门",
            "门按摩", "门保健", "蒙汗药", "孟建柱", "迷昏口", "迷奸", "迷魂香", "迷幻型", "迷幻药", "迷情药", "迷幻藥",
            "迷魂药", "迷藥", "迷奸药", "迷昏药", "迷昏藥", "迷魂藥", "迷情水", "谜奸药", "秘唇", "蜜穴", "密穴", "民抗议",
            "明慧网", "摸阴蒂", "某锦涛", "母奸", "母子乱伦", "母子奸情", "奶子", "男奴", "男女交欢", "内射", "嫩穴",
            "嫩b", "嫩逼", "嫩屄", "嫩bb", "嫩阴", "你妈死了", "你日妈", "你妈逼", "娘西皮", "娘了个比", "娘两腿之间",
            "浓精", "怒的志愿", "女优", "女任职名", "女人和狗", "女技师", "女激情", "女優", "女上门", "女被人家搞",
            "拍肩神药", "炮友", "喷尿", "屁眼", "平叫到床", "平惨案", "仆不怕饮", "普通嘌", "期货配", "奇迹的黄", "奇淫散",
            "强暴", "强硬发言", "强奸你妹", "强奸", "强权政府", "巧淫奸戏", "情色", "全裸", "全家死绝", "全家死光", "群奸暴",
            "群体性事", "群起抗暴", "群交", "绕过封锁", "人类灭亡", "人兽", "人妻", "人妻做爱", "人妻熟女", "人妻榨乳",
            "人体炸弹", "人妻色诱", "日中断交", "日你妹", "日烂", "日你全家", "日死你", "日你妈", "日逼", "肉棒", "肉壶",
            "肉蒲团", "肉便器", "肉茎", "肉淫器吞精", "肉棍", "肉穴", "肉逼", "肉棍干骚妇", "肉唇", "肉洞", "乳交",
            "软弱的国", "三秒倒", "三级片", "三挫", "三唑", "搔逼", "骚妇", "骚浪美女", "骚贱", "骚嘴", "骚姐姐",
            "骚母", "骚洞", "骚穴", "骚逼", "骚妹", "骚浪", "骚乳", "骚女", "骚水", "色妹妹", "色电影", "色视频",
            "色猫", "色小说", "色书库", "傻臂", "傻逼", "傻b", "傻避", "煞笔", "煞逼", "少妇", "舌头穴", "射颜",
            "射爽", "社会主义灭亡", "沈跃跃", "十八禁", "兽奸", "熟妇人妻", "爽片", "爽穴", "死全家", "死逼", "苏树林",
            "酥穴", "塔利班", "蘚鮑", "台湾", "台独", "台湾独立", "体奸", "天黯门", "天岸门", "天案们", "天安门大屠杀",
            "天氨门", "天安门", "天胺门", "天案门", "舔脚", "舔屄", "铁凝", "偷窥图片", "推背图", "退党", "吞精",
            "王胜俊", "王岐山", "王洛林", "王太华", "王鸿举", "瘟总理", "瘟加饱", "瘟假饱", "温影帝", "温家堡",
            "温切斯特", "温家宝", "温家某", "吴定富", "吴邦国", "吸精少女", "习近平", "洗肠射尿", "下流地带", "下贱",
            "销魂洞", "小穴", "小平遗言", "小逼崽子", "小平遗嘱", "新唐人", "性虐", "性奴", "性爱", "性爱图库",
            "颜射", "艳妇淫女", "要射精了", "要泄了", "要射了", "要人权", "耀邦", "叶剑英", "夜激情", "液体炸",
            "一夜欢", "遗嘱小平", "阴屄", "阴b", "阴茎", "阴唇", "阴部", "阴締", "淫贱", "淫叫", "淫液", "淫娃",
            "淫魔舞", "淫情", "淫肉", "淫蜜", "淫蕩", "淫汁", "淫姐", "淫妞", "淫逼", "淫亂潮吹", "淫荡", "淫騷妹",
            "淫情女", "淫奴", "淫兽", "淫兽学", "淫水", "淫妇", "淫穴", "淫水爱液", "淫母", "淫亂", "幼交", "玉穴",
            "原味内裤", "原味内衣", "援交妹", "杂种操的", "则民爷爷", "泽民", "炸立交", "炸弹遥控", "炸药的制",
            "炸药", "炸弹", "炸药配", "炸弹教", "炸鸟巢", "炸广州", "炸药制", "炸学校", "张志国", "张德江", "张荣坤",
            "政治局常委", "政fu", "政府", "制服狩", "中国人权", "中日断交", "朱镕基", "猪容鸡", "诸世纪", "主席",
            "主席像", "自焚", "自拍美穴", "自慰抠穴", "自慰", "总竖鸡", "总理", "总书记", "钻插", "做爱", "妞上门",
            "嫖妓指南", "嫖鸡", "嫖俄罗", "門服務", "陰道", "陰唇", "陰戶", "掰穴", "騷浪"
        ]

        // 构建DFA算法的词典树
        sensitiveWords.forEach(word => {
            let currentMap = this.sensitiveWordMap;
            for (let i = 0; i < word.length; i++) {
                const char = word.charAt(i);
                let wordMap = currentMap.get(char);

                if (wordMap) {
                    currentMap = wordMap;
                } else {
                    wordMap = new Map();
                    currentMap.set(char, wordMap);
                    currentMap = wordMap;
                }

                // 标记结束
                if (i === word.length - 1) {
                    currentMap.set('isEnd', true);
                }
            }
        });
    }

    // 检查文本是否包含敏感词
    public containsSensitiveWords(text: string): { found: boolean; words: string[] } {
        const foundWords: string[] = [];
        let currentMap: Map<string, any>;

        for (let i = 0; i < text.length; i++) {
            currentMap = this.sensitiveWordMap;
            let tempWord = '';

            for (let j = i; j < text.length; j++) {
                const char = text.charAt(j);
                const nextMap = currentMap.get(char);

                if (!nextMap) {
                    break;
                }

                tempWord += char;
                if (nextMap.get('isEnd')) {
                    foundWords.push(tempWord);
                    break;
                }

                currentMap = nextMap;
            }
        }

        return {
            found: foundWords.length > 0,
            words: foundWords
        };
    }

    // 替换敏感词为 *
    public filter(text: string): string {
        let result = text;
        const { words } = this.containsSensitiveWords(text);

        words.forEach(word => {
            const stars = '*'.repeat(word.length);
            result = result.replace(new RegExp(word, 'g'), stars);
        });

        return result;
    }
}

// 导出单例实例
export const sensitiveWordFilter = new SensitiveWordFilter(); 