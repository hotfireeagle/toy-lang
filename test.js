export const productTypes = new Map()
  .set(0, '税金贷')
  .set(1, '发票贷')
  .set(2, '经营贷')
  .set(3, '个贷类')
  .set(4, '车抵贷');

export const CREDIT_CARD_TYPE = 5;

export const cardProductTypes = new Map().set(CREDIT_CARD_TYPE, '信用卡');

export const allProductTypes = new Map()
  .set(0, '税金贷')
  .set(1, '发票贷')
  .set(2, '经营贷')
  .set(3, '个贷类')
  .set(4, '车抵贷')
  .set(CREDIT_CARD_TYPE, '信用卡');

export const DAI_TONGGUO = 3;
export const DAI_WEI_TONGGUO = 2;
export const DAI_YITIKUAN = 4;
export const XIN_BUTONGUO = 6;

export const progressTypes = new Map()
  .set(0, '贷款-已申请')
  .set(1, '贷款-授信中')
  .set(DAI_TONGGUO, '贷款-授信通过')
  .set(DAI_WEI_TONGGUO, '贷款-审批未通过')
  .set(DAI_YITIKUAN, '贷款-已提款')
  .set(5, '信用卡-待审核')
  .set(XIN_BUTONGUO, '信用卡-审核不通过')
  .set(7, '信用卡-审核已通过');

export const NOT_SETTLEMENT = 0;
export const YI_SHOU_XIN_JIE_SUAN = 1;
export const YI_JIE_SUAN = 2;

export const settlementStatusTypes = new Map()
  .set(NOT_SETTLEMENT, '未结算')
  .set(YI_SHOU_XIN_JIE_SUAN, '已授信结算')
  .set(YI_JIE_SUAN, '已结算');

export const offline = 0;
export const online = 1;

export const onlineStatusMap = new Map().set(offline, '下架').set(online, '上架');

export const repayTypeMap = new Map()
  .set('0', '等额本息')
  .set('1', '等额本金')
  .set('2', '等本等息')
  .set('3', '先息后本');

export const loanMonthMap = new Map()
  .set('3', '3期')
  .set('6', '6期')
  .set('9', '9期')
  .set('12', '12期')
  .set('15', '15期')
  .set('18', '18期')
  .set('24', '24期')
  .set('30', '30期')
  .set('36', '36期')
  .set('48', '48期')
  .set('60', '60期');

export const billType0 = 0;
export const billType1 = 1;
export const billType2 = 2;
export const billType3 = 3;
export const billType4 = 4;
export const billType5 = 5;
export const billType6 = 6;
export const billType7 = 7;

export const billTypeMap = new Map()
  .set(billType0, '授信结算')
  .set(billType1, '放款结算')
  .set(billType2, '组合结算')
  .set(billType3, '授信结算关联贷款周期')
  .set(billType4, '放款结算关联贷款周期')
  .set(billType5, '组合结算关联贷款周期')
  .set(billType6, '组合延期结算')
  .set(billType7, '放款延期结算');

export const allSettlementType = new Map()
  .set(billType0, '授信结算')
  .set(billType1, '放款结算')
  .set(billType2, '组合结算')
  .set(billType3, '授信结算关联贷款周期')
  .set(billType4, '放款结算关联贷款周期')
  .set(billType5, '组合结算关联贷款周期')
  .set(billType6, '组合延期结算')
  .set(billType7, '放款延期结算')
  .set(8, '首刷自动结算')
  .set(9, '核卡自动结算');

export const cardBillTypeMap = new Map().set(0, '首刷自动结算').set(1, '核卡自动结算');

export const adsTypeMap = new Map()
  .set(0, '首页左侧')
  .set(1, '首页右侧')
  .set(2, '我的中部')
  .set(3, '首页弹窗')
  .set(4, '产品大全弹窗')
  .set(5, '展业中心弹窗')
  .set(6, '我的弹窗')
  .set(7, '首页主推弹窗')
  .set(8, '我的积分');
// .set(9, '获客秘籍左侧')
// .set(10, '获客秘籍右侧');

export const settlementPeriodTypeMap = new Map().set(0, '日结');

export const statusMap = new Map()
  .set('40', '未划付')
  .set('20', '已划付')
  .set('30', '划付失败')
  .set('50', '划付中');

export const tixianShenPiStatusMap = new Map()
  .set('0', '待审核')
  .set('1', '已通过')
  .set('2', '已拒绝');

export const payTypeMap = new Map().set(0, '税务付费').set(1, '会员充值');

export const customerVipMap = new Map().set(1, '普通会员').set(2, '超级会员');

export const freezeTypeMap = new Map().set(1, '冻结').set(0, '解冻');

export const COMPANY_CUSTOMER_TYPE = 0;
export const customerTypeMap = new Map().set(COMPANY_CUSTOMER_TYPE, '企业客户').set(1, '个人客户');

export const managerLevelMap = new Map().set(1, '服务经理').set(2, '高级经理').set(3, '高级总监');

export const gongGaoStatusMap = new Map()
  .set(0, '未发布')
  .set(1, '已发布')
  .set(2, '已取消')
  .set(3, '发布失败')
  .set(4, '发布中');

export const ZHI_DING_USER = 3;

// 现金券的参与人群
export const ALL_CAN_CAN_YU = 0;
export const canYuRenQunMap = new Map()
  .set(ALL_CAN_CAN_YU, '全部')
  .set(1, '新用户')
  .set(2, '老用户')
  .set(ZHI_DING_USER, '指定用户');

export const couponStatusMap = new Map().set(0, '有效').set(1, '已失效');
export const couponTypeMap = new Map()
  .set(0, '进件现金券')
  .set(1, '邀请好友进件现金券')
  .set(2, '积分现金券');

export const sendTypeStatusMap = new Map().set(0, '正常').set(1, '停止发券');

export const userCouponStatusMap = new Map()
  .set(0, '待使用')
  .set(1, '即将过期')
  .set(2, '已使用')
  .set(3, '已失效');

export const couponProductType = new Map().set(0, '全部产品').set(1, '指定产品');

export const wxTemplateTypeMap = new Map().set(0, '已禁用').set(1, '已启用');

export const gatherStatusMap = new Map().set(0, '已下架').set(1, '已上架');
export const feedbackStatusMap = new Map()
  .set(0, '系统bug')
  .set(1, '优化体验')
  .set(2, '功能操作')
  .set(3, '账号登录')
  .set(5, '商务合作')
  .set(4, '其他问题');
export const taskStatusMap = new Map().set(0, '已下架').set(1, '已上架');
export const integraltatusMap = new Map().set(0, '正常').set(1, '已冻结');
export const goodsStatusMap = new Map().set(0, '已下架').set(1, '已上架');
export const goodsTypeMap = new Map().set(0, '现金兑换').set(1, '现金券兑换');

export const QUAN_HOU_TAI_ZENG_SONG = 0;
export const QUAN_SHOU_DONG_LING_QU = 1;

export const JI_FEN_PRODUCT_COUPON_TYPE = 2;

export const BUSINESS_CHANNEL_TYPE = 2;
export const channelTypeMap = new Map().set(1, '个人').set(BUSINESS_CHANNEL_TYPE, '企业');

// 渠道后台状态枚举值
export const BackgroundOpenStatus = 1; // 开放的状态枚举值
export const channelBackgroundColor = { 1: '#73d13d' };
export const channelBackgroundStatus = new Map().set(0, '已关闭').set(1, '已开放');

export const IncomingStatus = new Map().set(0, '未成交').set(1, '已成交');
export const incomingStatusColorObj = {
  0: 'red',
};

export const OfflineIncomingSource = 1;
export const IncomingSource = new Map().set(0, '线上').set(1, '线下');

export const CHANNEL_SCENE = 'channel_scene';
export const OPE_SCENE = 'ope_scene';

export const LABEL_QUESTION_TYPE = new Map().set('0', '单选').set('1', '多选');

export const LABEL_OFFLINE = 0;
export const LABEL_ONLINE = 1;
export const LABEL_STATUS_MAP = new Map().set(LABEL_ONLINE, '上架').set(LABEL_OFFLINE, '下架');

export const H5A_NOT_START = 0;
export const H5A_RUN = 1;
export const H5A_DONE = 2;
// export const H5A_CLOSE = 3;
export const H5A_BIANJI = 4;
// h5活动状态枚举值对象
export const H5ACTIVITY_STATUS = new Map()
  .set(H5A_NOT_START, '未开始')
  .set(H5A_RUN, '活动中')
  .set(H5A_DONE, '已结束')
  // .set(H5A_CLOSE, '已关闭')
  .set(H5A_BIANJI, '待编辑');

export const H5_ACTIVITITY_XIAJIA = 0;
export const H5_ACTIVITITY_SHANGJIA = 1;
export const H5_SHOW_STATUS = new Map()
  .set(H5_ACTIVITITY_XIAJIA, '下架')
  .set(H5_ACTIVITITY_SHANGJIA, '上架');

export const H5_LOTTERY = 0;
export const H5_INVITE = 1;
export const H5_SHARE = 2;
// h5活动类型枚举值对象
export const H5Activity_TYPE = new Map()
  .set(H5_LOTTERY, '抽奖')
  .set(H5_INVITE, '邀请')
  .set(H5_SHARE, '分享现金券活动');

export const H5_Lottery_CJ_TYPE_INVITE = 0;
export const H5_Lottery_CJ_TYPE_JIFEN = 1;
export const H5_Lottery_CJ_TYPE_CUSTOM = 2;

export const H5_Lottery_CJ_TYPE = new Map()
  .set(H5_Lottery_CJ_TYPE_INVITE, '邀请好友注册')
  .set(H5_Lottery_CJ_TYPE_JIFEN, '积分抽奖')
  .set(H5_Lottery_CJ_TYPE_CUSTOM, '自定义抽奖');

export const H5_CJ_LIMIT0 = 0;
export const H5_CJ_LIMIT1 = 1;

export const H5_CJ_LIMIT = new Map().set(H5_CJ_LIMIT0, '无限制').set(H5_CJ_LIMIT1, '有限制');

export const ZPNUM = new Map().set(6, '六等分').set(8, '八等分');

export const H5_REWARD_MONEY = 0;
export const H5_REWARD_QUAN = 1;
export const H5_REWARD_JIFEN = 2;
export const H5_REWARD_CUSTOM = 3;
export const H5_REWARD_THANK = 4;

export const REWARD = new Map()
  .set(H5_REWARD_MONEY, '现金红包')
  .set(H5_REWARD_QUAN, '现金券')
  .set(H5_REWARD_JIFEN, '积分')
  .set(H5_REWARD_CUSTOM, '自定义奖品')
  .set(H5_REWARD_THANK, '谢谢参与');

export const XIAN_JIN_QUAN_V1 = 0; // 自定义开始时间的枚举值
export const XIAN_JIN_QUAN_V2 = 1; // 领券当天为开始时间的枚举值
export const XIAN_JIN_QUAN_V3 = 2; // 领券次日为开始时间的枚举值

export const TEAM_COUNT = new Map().set(1, '上级').set(2, '上上级');

export const r0 = 'r0';
export const r1 = 'r1';
export const r2 = 'r2';
export const r3 = 'r3';

export const t0 = 0;
export const t1 = 1;
export const t2 = 2;
export const t3 = 3;

export const rotm = {
  [r0]: t0,
  [r1]: t1,
  [r2]: t2,
  [r3]: t3,
};

export const AS_RATE = 0; // 按比例的默认值
export const AS_AMOUNT = 1; // 按金额的默认值

export const contentTypeList = new Map().set(0, '图文').set(1, '视频');
export const content_status_daifabu = 0;
export const content_status_yifabu = 1;
export const content_status_yiquxiao = 2;
export const content_status_yixiajia = 3;

export const contentModuleStatusMap = new Map()
  .set(content_status_daifabu, '待发布 ')
  .set(content_status_yifabu, '已发布 ')
  .set(content_status_yiquxiao, '已取消')
  .set(content_status_yixiajia, '已下架');

export const position_zhaomu = 0;
export const position_huoke = 1;

export const MANUAL_WXMSG_WAIT = 0; // 待发送
export const MANUAL_WXMSG_SENDED = 1; // 已发送
export const MANUAL_WXMSG_CANCEL = 2; // 已取消
export const wxMsgStatusMap = new Map()
  .set(MANUAL_WXMSG_WAIT, '待发送')
  .set(MANUAL_WXMSG_SENDED, '已发送')
  .set(MANUAL_WXMSG_CANCEL, '已取消');

export const FISSION_NOT_START = 0;
export const FISSION_RUNNING = 1;
export const FISSION_END = 2;
export const FISSION_CLOSE = 3;
export const FISSION_EDIT = 4;

export const fissionStatusMap = new Map()
  .set(FISSION_NOT_START, '未开始')
  .set(FISSION_RUNNING, '活动中')
  .set(FISSION_END, '已结束')
  .set(FISSION_CLOSE, '已关闭')
  .set(FISSION_EDIT, '待编辑');

export const FISSION_SHOW_AVATAR = 1; // 显示头像昵称
export const FISSION_HIDE_AVATAR = 0; // 隐藏头像昵称

export const fissionShowAvatarMap = new Map()
  .set(FISSION_SHOW_AVATAR, '显示')
  .set(FISSION_HIDE_AVATAR, '隐藏');

export const FISSION_REDTYPE_RANDOM = 1;
export const FISSION_REDTYPE_FIXED = 2;
export const fissionRedtypeMap = new Map()
  .set(FISSION_REDTYPE_RANDOM, '随机')
  .set(FISSION_REDTYPE_FIXED, '固定');

export const FISSION_FINISH = 1;
export const FISSION_NOT_FINISH = 0;
export const fissionFinishMap = new Map()
  .set(FISSION_FINISH, '已完成')
  .set(FISSION_NOT_FINISH, '未完成');
