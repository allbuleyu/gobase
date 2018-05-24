package main

import (
	"github.com/allbuleyu/gobase/gopl.io/ch7"
	"golang.org/x/net/html"
	"fmt"
)


func main() {
	//resp, err := http.Get("http://www.baidu.com")
	//doc1, err := html.Parse(resp.Body)
	//fmt.Println(doc1, err)
	//return
	reader := ch7.NewReader(str)
	doc, err := html.Parse(reader)



	fmt.Println(doc, err)
}

var str = `
<!DOCTYPE html>
<!-- saved from url=(0042)http://dev.cms.catjc.com/index/index/index -->
<html><head><meta http-equiv="Content-Type" content="text/html; charset=UTF-8"><style class="vjs-styles-defaults">
      .video-js {
        width: 300px;
        height: 150px;
      }

      .video-js.vjs-fluid {
        padding-top: 56.25%
      }
    </style>



<title>竞彩猫后台管理系统</title>

<meta content="width=device-width, initial-scale=1.0,maximum-scale=1.0, user-scalable=0" name="viewport">
<meta http-equiv="X-UA-Compatible" content="IE=9;IE=8;IE=7;IE=EDGE">
<meta name="renderer" content="webkit|ie-comp|ie-stand">

<meta content="" name="description">
<meta content="" name="author">

<!-- BEGIN GLOBAL MANDATORY STYLES -->

<link href="./xxx_files/bootstrap.min.css" rel="stylesheet" type="text/css">

<link href="./xxx_files/bootstrap-responsive.min.css" rel="stylesheet" type="text/css">

<link href="./xxx_files/font-awesome.min.css" rel="stylesheet" type="text/css">

<link href="./xxx_files/glyphicons.css" rel="stylesheet" type="text/css">

<link href="./xxx_files/style-metro.css" rel="stylesheet" type="text/css">

<link href="./xxx_files/style.css" rel="stylesheet" type="text/css">

<link href="./xxx_files/style-responsive.css" rel="stylesheet" type="text/css">

<link href="./xxx_files/default.css" rel="stylesheet" type="text/css" id="style_color">
<link href="./xxx_files/uniform.default.css" rel="stylesheet" type="text/css">
<!-- END GLOBAL MANDATORY STYLES -->
<!-- BEGIN PAGE LEVEL STYLES -->
<link href="./xxx_files/jquery.gritter.css" rel="stylesheet" type="text/css">

<link href="./xxx_files/daterangepicker.css" rel="stylesheet" type="text/css">

<link href="./xxx_files/fullcalendar.css" rel="stylesheet" type="text/css">

<link href="./xxx_files/jqvmap.css" rel="stylesheet" type="text/css" media="screen">

<link href="./xxx_files/jquery.easy-pie-chart.css" rel="stylesheet" type="text/css" media="screen">

<link rel="stylesheet" type="text/css" href="./xxx_files/datepicker.css">

<link rel="stylesheet" type="text/css" href="./xxx_files/datetimepicker.css">

<link href="./xxx_files/jquery.fancybox.css" rel="stylesheet">

<!-- <link rel="stylesheet" type="text/css" href="/resource/media/css/select2_metro.css" /> -->

<link rel="stylesheet" href="./xxx_files/DT_bootstrap.css">

<!-- <link href="/resource/media/css/dropzone.css" rel="stylesheet"/> -->

<!-- END PAGE LEVEL STYLES -->
<link rel="shortcut icon" href="http://dev.cms.catjc.com/resource/media/image/favicon.ico">

<!-- custom-->
<link rel="stylesheet" href="./xxx_files/my-style.css">


<!-- plugin -->
<link href="./xxx_files/jquery.fileupload.css" rel="stylesheet" type="text/css" id="style_color">
<!-- <link rel="stylesheet" type="text/css" href="/resource/plugin/fex-team-webuploader/css/webuploader.css" /> -->

<!-- <link rel="stylesheet" type="text/css" href="/resource/media/css/bootstrap-fileupload.css" /> -->

<link href="./xxx_files/jquery.fileupload-ui.css" rel="stylesheet">
<!--<link href="/resource/media/css/nep.min.css" rel="stylesheet" type="text/css"/>-->
<link href="./xxx_files/nep.min.css" rel="stylesheet">
<style>
    a {
        color: #737373;
        text-decoration: none;
    }
    #icon{
        font-size: 18px;
    }
    .scroll-wrapper {
        -webkit-overflow-scrolling: touch ! important;
        overflow-y: scroll ! important;
    }
    @media (max-width:979px){
        #div-page-sidebar{
            display: none;
        }
        #table_fb,#table_r9,#table_bd{
            width:600px;
            border-collapse: collapse;
            background: #fff;
        }
        #table_fb .fb_spf{
            width:100px;
        }
        #table_fb .fb_rq{
            width:30px;
        }
        #table_fb .fb_sfp{
            width:150px;
        }
        #table_bd .bd_xx{
            width: 80px;
        }

    }
    .closebtn{
        background:#333;
        font-size: 12px;
        position:absolute;
        right:-8px;
        top:-8px;
        border-radius:50%;
        width:20px;
        height:20px;
        line-height:20px;
        border: solid 1px #FFF;
        text-align: center;
        color: #FFF;
        z-index: 999999999;
    }

</style>




	<link rel="stylesheet" href="./xxx_files/layer.css" id="layuicss-skinlayercss"><style type="text/css">.jqstooltip { position: absolute;left: 0px;top: 0px;visibility: hidden;background: rgb(0, 0, 0) transparent;background-color: rgba(0,0,0,0.6);filter:progid:DXImageTransform.Microsoft.gradient(startColorstr=#99000000, endColorstr=#99000000);-ms-filter: "progid:DXImageTransform.Microsoft.gradient(startColorstr=#99000000, endColorstr=#99000000)";color: white;font: 10px arial, san serif;text-align: left;white-space: nowrap;padding: 5px;border: 1px solid white;z-index: 10000;}.jqsfield { color: white;font: 10px arial, san serif;text-align: left;}</style></head>
	<body class="page-header-fixed" style="">

			<!-- begin header -->
<div class="header navbar navbar-inverse navbar-fixed-top">
	<div class="navbar-inner" style="height: 100%">
		<div class="container-fluid">
			<a class="brand" href="http://dev.cms.catjc.com/index/index/index.html" style="padding: 10px 0 "><img src="./xxx_files/logo.png" alt="logo"></a>
			<div class="dropdown hidden-desktop" style="float: right;padding: 10px 0;">
				<a href="http://dev.cms.catjc.com/index/index/index###" class="dropdown-toggle" data-toggle="dropdown" onclick="hiddenSee()">
					<span id="icon"><i class="icon-align-justify"></i></span>
				</a>
			</div>
			<div style="display: inline-block;float: right">
				<ul class="nav pull-right">
					<!--<li class="dropdown hidden-desktop">-->
						<!---->

					<!--</li>-->
					<li class="dropdown" id="hlader_inbox_bar">
						<a href="http://dev.cms.catjc.com/index/index/index#" class="dropdown-toggle" data-toggle="dropdown">
							<i class="icon-envelope"></i>
							<span class="badge">0</span>
						</a>
						<ul class="dropdown-menu extended inbox">
							<!-- <li><p>5 条推荐待审核</p></li><li>
							<li>
								<a href="javascript:showpop('推荐审核','../push/pushcheck?id=96472')">
									<span class="photo">
										<img src="/public/uploads/expert/1456987682.png" alt="">
									</span>
									<span class="subject"><span class="from">
										Brothers
									</span>
									<span class="time">1030小时</span>
									</span><span class="message">asdfasdf</span>
								</a>
							</li>
							<li class="external"><a href="../push/pushchecklist">查看全部 <i class="m-icon-swapright"></i></a></li> -->
						</ul>
					</li>
					<li class="dropdown user">
						<a href="http://dev.cms.catjc.com/index/index/index#" class="dropdown-toggle" data-toggle="dropdown">
							<img alt="" src="./xxx_files/avatar1_small.jpg">
							<span class="username">阿黄666</span>
							<i class="icon-angle-down"></i>
						</a>
						<ul class="dropdown-menu">
							 <!--<li><a href="extra_lock.html"><i class="icon-lock"></i>修改密码</a></li>-->
							<li><a href="http://dev.cms.catjc.com/auth/login/logout.html"><i class="icon-key"></i>退出</a></li>
						</ul>
					</li>
				</ul>
			</div>
		</div>
	</div>
</div>
<!-- end header -->




		<div class="page-container row-fluid">

				<!--begin sidebar left-->
<div class="page-sidebar" id="div-page-sidebar">
	<!-- BEGIN SIDEBAR MENU -->
	<ul class="page-sidebar-menu">
		<li>
			<!-- BEGIN SIDEBAR TOGGLER BUTTON -->
			<div class="sidebar-toggler hidden-phone"></div>
			<!-- BEGIN SIDEBAR TOGGLER BUTTON -->
		</li>

		<li class="start active">
			<a href="http://dev.cms.catjc.com/"><i class="icon-home"></i><span class="title">首页</span><span class="selected"></span></a>
		</li>



				<li class="start ">
			<a href="javascript:;"><i class="icon-table"></i><span class="title">系统管理</span><span class="arrow "></span></a>
			<ul class="sub-menu">

									<li><a href="http://dev.cms.catjc.com/sys/module/moduleList.html">模块管理</a></li>
										<li><a href="http://dev.cms.catjc.com/sys/role/roleList.html">角色管理</a></li>
										<li><a href="http://dev.cms.catjc.com/sys/user/userList.html">系统用户管理</a></li>
										<li><a href="http://dev.cms.catjc.com/sys/cache/cachelist.html">缓存更新</a></li>
										<li><a href="http://dev.cms.catjc.com/sys/configure/adconfigurelist.html">广告配置项管理</a></li>
										<li><a href="http://dev.cms.catjc.com/sys/configure/mainconfigure.html">主配置管理</a></li>
										<li><a href="http://dev.cms.catjc.com/sys/configure/topBanner.html">栏目名管理</a></li>
										<li><a href="http://dev.cms.catjc.com/sys/configure/navigation.html">导航管理栏目</a></li>
										<li><a href="http://dev.cms.catjc.com/sys/configure/words.html">蔽字库管理</a></li>
										<li><a href="http://dev.cms.catjc.com/sys/coo/cooList.html">渠道用户管理</a></li>
								</ul>
		</li>
				<li class="start ">
			<a href="javascript:;"><i class="icon-table"></i><span class="title">用户管理</span><span class="arrow "></span></a>
			<ul class="sub-menu">

									<li><a href="http://dev.cms.catjc.com/user/user/userList.html">用户列表</a></li>
										<li><a href="http://dev.cms.catjc.com/user/user/headCheckList.html">头像审核</a></li>
										<li><a href="http://dev.cms.catjc.com/user/user/nicknameCheckList.html">昵称审核</a></li>
										<li><a href="http://dev.cms.catjc.com/user/usermember/userMemberList.html">用户会员</a></li>
										<li><a href="http://dev.cms.catjc.com/user/user/userListAll.html">用户列表(全部)</a></li>
										<li><a href="http://dev.cms.catjc.com/user/user/autoVip.html">用户会员(自动)</a></li>
										<li><a href="http://dev.cms.catjc.com/operation/finance/vipPostPone.html">会员延期</a></li>
										<li><a href="http://dev.cms.catjc.com/operation/finance/postponeLog.html">延期日志</a></li>
								</ul>
		</li>
				<li class="start ">
			<a href="javascript:;"><i class="icon-table"></i><span class="title">推荐管理</span><span class="arrow "></span></a>
			<ul class="sub-menu">

									<li><a href="http://dev.cms.catjc.com/recommend/jc/jcList.html">竞彩推荐</a></li>
										<li><a href="http://dev.cms.catjc.com/recommend/match/matchList.html">赛事推荐</a></li>
										<li><a href="http://dev.cms.catjc.com/recommend/audit/auditList.html">审核推荐</a></li>
										<li><a href="http://dev.cms.catjc.com/recommend/fp/expList.html">投注经验</a></li>
										<li><a href="http://dev.cms.catjc.com/recommend/result/resultlist.html">成绩明细</a></li>
										<li><a href="http://dev.cms.catjc.com/recommend/result/resultCollectList.html">成绩汇总</a></li>
										<li><a href="http://dev.cms.catjc.com/recommend/video/videoList.html">视频推荐</a></li>
										<li><a href="http://dev.cms.catjc.com/recommend/video/auditList.html">视频审核</a></li>
										<li><a href="http://dev.cms.catjc.com/recommend/match/expertMarkList.html">专家总结统计</a></li>
								</ul>
		</li>
				<li class="start ">
			<a href="javascript:;"><i class="icon-table"></i><span class="title">内容管理</span><span class="arrow "></span></a>
			<ul class="sub-menu">

									<li><a href="http://dev.cms.catjc.com/content/head/headList.html">竞彩头条</a></li>
										<li><a href="http://dev.cms.catjc.com/content/summary/summaryList.html">专家总结</a></li>
										<li><a href="http://dev.cms.catjc.com/content/single/singleList.html">单页内容管理</a></li>
										<li><a href="http://dev.cms.catjc.com/content/ad/adList.html">广告管理</a></li>
										<li><a href="http://dev.cms.catjc.com/content/slide/slideList.html">轮播图管理</a></li>
										<li><a href="http://dev.cms.catjc.com/content/help/helpList.html">使用帮助</a></li>
										<li><a href="http://dev.cms.catjc.com/content/leitai/leitaiList.html">擂台轮播图</a></li>
										<li><a href="http://dev.cms.catjc.com/content/leitai/leitaiPeriodList.html">擂台管理</a></li>
										<li><a href="http://dev.cms.catjc.com/content/leitai/leitaiResult.html">任九擂台榜单</a></li>
										<li><a href="http://dev.cms.catjc.com/content/leitai/soloResult.html">单刀赴会榜单</a></li>
										<li><a href="http://dev.cms.catjc.com/content/solo/noticeList.html">单刀赴会公告管理</a></li>
										<li><a href="http://dev.cms.catjc.com/recommend/match/luckStarList.html">幸运星统计</a></li>
										<li><a href="http://dev.cms.catjc.com/content/cup/cupList.html">2018世界杯专区</a></li>
										<li><a href="http://dev.cms.catjc.com/content/cup/sortTeam.html">世界杯球队</a></li>
										<li><a href="http://dev.cms.catjc.com/content/solo/totalIntegral.html">总积分榜</a></li>
										<li><a href="http://dev.cms.catjc.com/content/solo/soloThirtyTwo.html">32强录入</a></li>
										<li><a href="http://dev.cms.catjc.com/content/solo/activitySettle.html">单刀活动榜单</a></li>
										<li><a href="http://dev.cms.catjc.com/content/slide/ttt.html">轮播图管理</a></li>
								</ul>
		</li>
				<li class="start ">
			<a href="javascript:;"><i class="icon-table"></i><span class="title">赛事管理</span><span class="arrow "></span></a>
			<ul class="sub-menu">

									<li><a href="http://dev.cms.catjc.com/match/team/bbTeamList.html">篮球队图标</a></li>
										<li><a href="http://dev.cms.catjc.com/match/team/fbTeamList.html">足球队图标</a></li>
										<li><a href="http://dev.cms.catjc.com/match/bind/matchList.html">赛事绑定</a></li>
										<li><a href="http://dev.cms.catjc.com/match/bind/teamList.html">球队绑定</a></li>
										<li><a href="http://dev.cms.catjc.com/match/fbmatch/fbmatchlist.html">足球赛事管理</a></li>
										<li><a href="http://dev.cms.catjc.com/match/bbmatch/bbmatchlist.html">篮球赛事管理</a></li>
										<li><a href="http://dev.cms.catjc.com/match/fbmatch/score.html">足球比分监控</a></li>
										<li><a href="http://dev.cms.catjc.com/match/fbmatch/hot.html">热门赛事</a></li>
										<li><a href="http://dev.cms.catjc.com/match/fbmatch/ds_score.html">DS足球监控</a></li>
										<li><a href="http://dev.cms.catjc.com/match/fbmatch/europe.html">欧冠闯关</a></li>
								</ul>
		</li>
				<li class="start ">
			<a href="javascript:;"><i class="icon-table"></i><span class="title">专家管理</span><span class="arrow "></span></a>
			<ul class="sub-menu">

									<li><a href="http://dev.cms.catjc.com/expert/expert/expertList.html">专家信息</a></li>
										<li><a href="http://dev.cms.catjc.com/expert/expert/sentList.html">推荐分发</a></li>
										<li><a href="http://dev.cms.catjc.com/expert/sales/sales_list.html">专家销量管理</a></li>
								</ul>
		</li>
				<li class="start ">
			<a href="javascript:;"><i class="icon-table"></i><span class="title">消息管理</span><span class="arrow "></span></a>
			<ul class="sub-menu">

									<li><a href="http://dev.cms.catjc.com/message/push/pushlist.html">推送管理</a></li>
										<li><a href="http://dev.cms.catjc.com/message/feedback/feedbacklist.html">意见反馈</a></li>
										<li><a href="http://dev.cms.catjc.com/message/message/msgList.html">发送日志</a></li>
								</ul>
		</li>
				<li class="start ">
			<a href="javascript:;"><i class="icon-table"></i><span class="title">数据分析</span><span class="arrow "></span></a>
			<ul class="sub-menu">

									<li><a href="http://dev.cms.catjc.com/dataanalysis/phplog/mapilist.html">接口数据统计</a></li>
										<li><a href="http://dev.cms.catjc.com/dataanalysis/phplog/ad.html">广告访问量</a></li>
										<li><a href="http://dev.cms.catjc.com/dataanalysis/phplog/scan.html">二维码扫描量</a></li>
										<li><a href="http://dev.cms.catjc.com/dataanalysis/phplog/soloData.html">单刀数据分析</a></li>
								</ul>
		</li>
				<li class="start ">
			<a href="javascript:;"><i class="icon-table"></i><span class="title">财务管理</span><span class="arrow "></span></a>
			<ul class="sub-menu">

									<li><a href="http://dev.cms.catjc.com/finance/finance/all.html">资金总览</a></li>
										<li><a href="http://dev.cms.catjc.com/finance/channel/all.html">渠道充值总揽</a></li>
										<li><a href="http://dev.cms.catjc.com/finance/expert/all.html">专家销售额总揽</a></li>
										<li><a href="http://dev.cms.catjc.com/finance/refund/refundlist.html">退款管理(客服/运营)</a></li>
										<li><a href="http://dev.cms.catjc.com/finance/finance/consumeService.html">用户消费查询(客服)</a></li>
										<li><a href="http://dev.cms.catjc.com/finance/finance/incomerecordlist.html">充值记录(客服)</a></li>
										<li><a href="http://dev.cms.catjc.com/finance/finance/incomerecordlistall.html">充值记录(全部)</a></li>
										<li><a href="http://dev.cms.catjc.com/finance/finance/consumeServiceall.html">用户消费查询(全部)</a></li>
								</ul>
		</li>
				<li class="start ">
			<a href="javascript:;"><i class="icon-table"></i><span class="title">运营对账分析</span><span class="arrow "></span></a>
			<ul class="sub-menu">

									<li><a href="http://dev.cms.catjc.com/operation/finance/incomeRecordList.html">第三方充值</a></li>
										<li><a href="http://dev.cms.catjc.com/operation/finance/detail.html">资金明细</a></li>
										<li><a href="http://dev.cms.catjc.com/operation/finance/consume.html">消费 查询</a></li>
										<li><a href="http://dev.cms.catjc.com/operation/finance/stream.html">购买流水</a></li>
										<li><a href="http://dev.cms.catjc.com/operation/finance/refundToFish.html">退款转鱼</a></li>
										<li><a href="http://dev.cms.catjc.com/operation/finance/salenumber.html">销售统计</a></li>
										<li><a href="http://dev.cms.catjc.com/operation/finance/deduct.html">手动扣款</a></li>
										<li><a href="http://dev.cms.catjc.com/operation/finance/activity.html">手动补单</a></li>
										<li><a href="http://dev.cms.catjc.com/operation/finance/solo.html">单刀赴会统计</a></li>
										<li><a href="http://dev.cms.catjc.com/operation/finance/alipayActivity.html">支付宝手动补单</a></li>
										<li><a href="http://dev.cms.catjc.com/operation/finance/channelStatistics.html">渠道统计</a></li>
										<li><a href="http://dev.cms.catjc.com/operation/finance/actChannelStatistics.html">活动渠道统计</a></li>
								</ul>
		</li>
				<li class="start ">
			<a href="javascript:;"><i class="icon-table"></i><span class="title">合作分析</span><span class="arrow "></span></a>
			<ul class="sub-menu">

									<li><a href="http://dev.cms.catjc.com/cooperate/cooperate/buywaterlist.html">购买明细</a></li>
										<li><a href="http://dev.cms.catjc.com/cooperate/cooperate/expertanalysis.html">专家分析</a></li>
										<li><a href="http://dev.cms.catjc.com/cooperate/cooperate/recommend.html">推荐总销量</a></li>
										<li><a href="http://dev.cms.catjc.com/cooperate/cooperate/saleanalysis.html">销量分析</a></li>
								</ul>
		</li>
				<li class="start ">
			<a href="javascript:;"><i class="icon-table"></i><span class="title">直营店管理</span><span class="arrow "></span></a>
			<ul class="sub-menu">

									<li><a href="http://dev.cms.catjc.com/store/store/dailyJoinList.html">日销售交接表</a></li>
										<li><a href="http://dev.cms.catjc.com/store/store/dailyDetailList.html">日销量详细表</a></li>
										<li><a href="http://dev.cms.catjc.com/store/store/dailyJoinInput.html">日销售交接表(录入)</a></li>
										<li><a href="http://dev.cms.catjc.com/store/store/dailyDetailInput.html">日销量详细表(录入)</a></li>
								</ul>
		</li>
				<li class="start ">
			<a href="javascript:;"><i class="icon-table"></i><span class="title">视频管理</span><span class="arrow "></span></a>
			<ul class="sub-menu">

									<li><a href="http://dev.cms.catjc.com/video/video/videolist.html">节目单管理</a></li>
								</ul>
		</li>
				<li class="start ">
			<a href="javascript:;"><i class="icon-table"></i><span class="title">年会相关</span><span class="arrow "></span></a>
			<ul class="sub-menu">

									<li><a href="http://dev.cms.catjc.com/annual/index/raffleList.html">人员列表</a></li>
								</ul>
		</li>
				<li class="start ">
			<a href="javascript:;"><i class="icon-table"></i><span class="title">数据分析(尽调)</span><span class="arrow "></span></a>
			<ul class="sub-menu">

									<li><a href="http://dev.cms.catjc.com/jindiao/index/expert.html">专家销量</a></li>
										<li><a href="http://dev.cms.catjc.com/jindiao/index/user.html">注册用户</a></li>
										<li><a href="http://dev.cms.catjc.com/jindiao/index/recharge.html">充值用户</a></li>
										<li><a href="http://dev.cms.catjc.com/jindiao/index/buy.html">消费用户</a></li>
										<li><a href="http://dev.cms.catjc.com/jindiao/index/dau.html">日总计</a></li>
										<li><a href="http://dev.cms.catjc.com/jindiao/index/mau.html">月总计</a></li>
										<li><a href="http://dev.cms.catjc.com/jindiao/index/userAttr.html">用户属性</a></li>
										<li><a href="http://dev.cms.catjc.com/jindiao/index/province.html">地域分布</a></li>
										<li><a href="http://dev.cms.catjc.com/jindiao/index/userRetention.html">用户留存</a></li>
								</ul>
		</li>
				<li class="start ">
			<a href="javascript:;"><i class="icon-table"></i><span class="title">单刀赴会</span><span class="arrow "></span></a>
			<ul class="sub-menu">

									<li><a href="http://dev.cms.catjc.com/solo/index/realTime.html">单刀实时数据</a></li>
										<li><a href="http://dev.cms.catjc.com/solo/index/data.html">单刀数据</a></li>
										<li><a href="http://dev.cms.catjc.com/solo/index/hits.html">单刀成绩</a></li>
								</ul>
		</li>
				<li class="start ">
			<a href="javascript:;"><i class="icon-table"></i><span class="title">官网管理</span><span class="arrow "></span></a>
			<ul class="sub-menu">

									<li><a href="http://dev.cms.catjc.com/www/classroom/classroomList.html">竞彩课堂</a></li>
										<li><a href="http://dev.cms.catjc.com/www/slide/slideList.html">专家轮播图</a></li>
										<li><a href="http://dev.cms.catjc.com/www/partner/partnerList.html">合作伙伴</a></li>
										<li><a href="http://dev.cms.catjc.com/www/expert/hotExpertList.html">热门专家</a></li>
										<li><a href="http://dev.cms.catjc.com/www/expert/hotAnchorList.html">主播推荐</a></li>
										<li><a href="http://dev.cms.catjc.com/www/officialslide/slideList.html">官网轮播图</a></li>
										<li><a href="http://dev.cms.catjc.com/www/configure/ADConfigureList.html">官网广告管理</a></li>
										<li><a href="http://dev.cms.catjc.com/www/foucs/foucsList.html">焦点资讯配置</a></li>
										<li><a href="http://dev.cms.catjc.com/www/info/infoList.html">资讯文章管理</a></li>
								</ul>
		</li>
				<li class="start ">
			<a href="javascript:;"><i class="icon-table"></i><span class="title">竞猜猫管理</span><span class="arrow "></span></a>
			<ul class="sub-menu">

									<li><a href="http://dev.cms.catjc.com/essence/essence/essenceArticleList.html">文章管理</a></li>
										<li><a href="http://dev.cms.catjc.com/essence/essence/videoList.html">视频管理</a></li>
										<li><a href="http://dev.cms.catjc.com/essence/essence/picList.html">图片管理</a></li>
										<li><a href="http://dev.cms.catjc.com/essence/questionbank/questionList.html">题库管理</a></li>
										<li><a href="http://dev.cms.catjc.com/essence/slide/slideList.html">轮播图</a></li>
								</ul>
		</li>
				<li class="start ">
			<a href="javascript:;"><i class="icon-table"></i><span class="title">竞彩猫大数据</span><span class="arrow "></span></a>
			<ul class="sub-menu">

									<li><a href="http://dev.cms.catjc.com/bigdata/user/userData.html">用户数据</a></li>
										<li><a href="http://dev.cms.catjc.com/bigdata/user/rechargeData.html">充值数据</a></li>
										<li><a href="http://dev.cms.catjc.com/bigdata/user/consumeData.html">消费数据</a></li>
										<li><a href="http://dev.cms.catjc.com/bigdata/expert/expertAllData.html">专家数据</a></li>
										<li><a href="http://dev.cms.catjc.com/bigdata/user/productData.html">包月产品</a></li>
										<li><a href="http://dev.cms.catjc.com/bigdata/user/hourData.html">时段数据</a></li>
										<li><a href="http://dev.cms.catjc.com/bigdata/user/detailData.html">资金流水</a></li>
										<li><a href="http://dev.cms.catjc.com/bigdata/replenish/day.html">日补充数据</a></li>
										<li><a href="http://dev.cms.catjc.com/bigdata/replenish/month.html">月补充数据</a></li>
								</ul>
		</li>
				<li class="start ">
			<a href="javascript:;"><i class="icon-table"></i><span class="title">竞彩公开课管理</span><span class="arrow "></span></a>
			<ul class="sub-menu">

									<li><a href="http://dev.cms.catjc.com/recommend/Jcvideo/JcVideoList.html">视频上传</a></li>
										<li><a href="http://dev.cms.catjc.com/ccm/slide/slideList.html">轮播图</a></li>
								</ul>
		</li>
				<li class="start ">
			<a href="javascript:;"><i class="icon-table"></i><span class="title">聊比赛管理</span><span class="arrow "></span></a>
			<ul class="sub-menu">

									<li><a href="http://dev.cms.catjc.com/recommend/Talkschedule/JcVideoList.html">聊比赛视频上传</a></li>
										<li><a href="http://dev.cms.catjc.com/ccm/Talkjcslide/slideList.html">聊比赛轮播图</a></li>
								</ul>
		</li>
				<li class="start ">
			<a href="javascript:;"><i class="icon-table"></i><span class="title">专题页管理</span><span class="arrow "></span></a>
			<ul class="sub-menu">

									<li><a href="http://dev.cms.catjc.com/special/Index/scoreboardSetting.html">积分榜设置</a></li>
										<li><a href="http://dev.cms.catjc.com/special/Focusinformation/focusList.html">焦点资讯</a></li>
								</ul>
		</li>
			</ul>
</div>
<!-- END sidebar left -->

			<div class="page-content" style="min-height:1008px !important">
				<div class="container-fluid">
					<div class="row-fluid">
						<div class="span12">
							<h3 class="page-title">	快捷导航</h3>
							<ul class="breadcrumb">
								<li>
									<i class="icon-home"></i>
									<a href="http://dev.cms.catjc.com/index/index/index.html">首页</a>
									<i class="icon-angle-right"></i>
								</li>
								<li><a href="http://dev.cms.catjc.com/index/index/index#">快捷导航</a></li>


							</ul>

							<!-- END PAGE TITLE & BREADCRUMB-->

						</div>



					</div>

	<!-- BEGIN PAGE CONTENT-->

	<div class="row-fluid">


		<div class="span12">

			<!-- BEGIN EXAMPLE TABLE PORTLET-->

			<div class="portlet box light-grey" style="border-top: 1px solid #eee;">


				<div class="portlet-body">

					<div class="clearfix">
						<div class="span3">

						</div>
						<div class="span6 text-center">
							<img src="./xxx_files/index.png" alt="">
							<p class="text-center">欢迎进入卡特猫管理系统</p>
						</div>
					</div>
				</div>

			</div>

			<!-- END EXAMPLE TABLE PORTLET-->

		</div>

	</div>


				</div>
			</div>

		</div>

			<div class="footer">
	<div class="footer-inner">2017 © catjc.com 竞彩猫</div>
	<div class="footer-tools"><span class="go-top"><i class="icon-angle-up"></i></span></div>
</div>


			<!-- END FOOTER -->
<!-- BEGIN JAVASCRIPTS(Load javascripts at bottom, this will reduce page load time) -->
<!-- BEGIN CORE PLUGINS -->
<script src="./xxx_files/jquery-1.10.1.min.js" type="text/javascript"></script>
<script src="./xxx_files/jquery-migrate-1.2.1.min.js" type="text/javascript"></script>

<!-- IMPORTANT! Load jquery-ui-1.10.1.custom.min.js before bootstrap.min.js to fix bootstrap tooltip conflict with jquery ui tooltip -->

<script src="./xxx_files/jquery-ui-1.10.1.custom.min.js" type="text/javascript"></script>
<script src="./xxx_files/bootstrap.min.js" type="text/javascript"></script>

<!--[if lt IE 9]>
<script src="/resource/media/js/excanvas.min.js"></script>
<script src="/resource/media/js/respond.min.js"></script>
<![endif]-->

<script src="./xxx_files/jquery.slimscroll.min.js" type="text/javascript"></script>

<script src="./xxx_files/jquery.blockui.min.js" type="text/javascript"></script>

<script src="./xxx_files/jquery.cookie.min.js" type="text/javascript"></script>

<!-- 改变checkbox样式的js -->
<script src="./xxx_files/jquery.uniform.min.js" type="text/javascript"></script>

<!-- END CORE PLUGINS -->

<!-- BEGIN PAGE LEVEL PLUGINS -->

<script src="./xxx_files/jquery.vmap.js" type="text/javascript"></script>

<script src="./xxx_files/jquery.vmap.russia.js" type="text/javascript"></script>

<script src="./xxx_files/jquery.vmap.world.js" type="text/javascript"></script>

<script src="./xxx_files/jquery.vmap.europe.js" type="text/javascript"></script>

<script src="./xxx_files/jquery.vmap.germany.js" type="text/javascript"></script>

<script src="./xxx_files/jquery.vmap.usa.js" type="text/javascript"></script>

<script src="./xxx_files/jquery.vmap.sampledata.js" type="text/javascript"></script>

<script src="./xxx_files/jquery.flot.js" type="text/javascript"></script>

<script src="./xxx_files/jquery.flot.resize.js" type="text/javascript"></script>

<script src="./xxx_files/jquery.pulsate.min.js" type="text/javascript"></script>

<script src="./xxx_files/date.js" type="text/javascript"></script>

<script src="./xxx_files/daterangepicker.js" type="text/javascript"></script>

<script src="./xxx_files/jquery.gritter.js" type="text/javascript"></script>

<script src="./xxx_files/fullcalendar.min.js" type="text/javascript"></script>

<script src="./xxx_files/jquery.easy-pie-chart.js" type="text/javascript"></script>

<script src="./xxx_files/jquery.sparkline.min.js" type="text/javascript"></script>

<script type="text/javascript" src="./xxx_files/bootstrap-datepicker.js"></script>

<script type="text/javascript" src="./xxx_files/bootstrap-datetimepicker.js"></script>

<script type="text/javascript" src="./xxx_files/bootstrap-datepicker.zh-CN.min.js"></script>

<script type="text/javascript" src="./xxx_files/bootstrap-datetimepicker.zh-CN.js"></script>


<!-- END PAGE LEVEL PLUGINS -->

<!-- BEGIN PAGE LEVEL SCRIPTS -->

<script src="./xxx_files/app.js" type="text/javascript"></script>

<script src="./xxx_files/index.js" type="text/javascript"></script>


<script type="text/javascript" src="./xxx_files/select2.min.js"></script>

<script type="text/javascript" src="./xxx_files/jquery.dataTables.js"></script>

<script type="text/javascript" src="./xxx_files/DT_bootstrap.js"></script>

<script src="./xxx_files/form-components.js"></script>

<!-- END PAGE LEVEL SCRIPTS -->

<script src="./xxx_files/table-managed.js"></script>

<!-- <script src="/resource/media/js/dropzone.js"></script> -->

<!-- plugin -->
<!-- <script src="/resource/plugin/fex-team-webuploader/dist/webuploader.min.js"></script>  -->

<script src="./xxx_files/jquery.ui.widget.js"></script>
<script src="./xxx_files/jquery.fileupload.js"></script>

<!--弹出插件layer-->
<script src="./xxx_files/layer.js" type="text/javascript"></script>
<!--<script src="/resource/components/layermobile/layerm.js" type="text/javascript"></script>-->
<!-- <script src="/resource/components/layer/layer.js" type="text/javascript"></script> -->

<script src="./xxx_files/custom.js" type="text/javascript"></script>
<!-- END 弹出插件layer-->

<script type="text/javascript" src="./xxx_files/jquery.validate.min.js"></script>

<script type="text/javascript" src="./xxx_files/additional-methods.min.js"></script>

<script type="text/javascript" src="./xxx_files/chosen.jquery.min.js"></script>


<!-- 编辑器 -->
<script type="text/javascript" charset="utf-8" src="./xxx_files/ueditor.config.js"></script>

<script type="text/javascript" charset="utf-8" src="./xxx_files/ueditor.all.min.js"> </script>
<!--bilibili.flv-->
<script type="text/javascript" charset="utf-8" src="./xxx_files/flv.min.js"> </script>
<!--网易云flv.js-->
<!--<script type="text/javascript" src="/resource/media/js/nep.min.js"></script>-->

<script src="./xxx_files/nep.min.js"></script>


<script>

	jQuery(document).ready(function() {

	   App.init(); // initlayout and core plugins

	   // Index.init();

	   // Index.initJQVMAP(); // init index page's custom scripts

	   // Index.initCalendar(); // init index page's custom scripts

	   // Index.initCharts(); // init index page's custom scripts

	   // Index.initChat();

	   // Index.initMiniCharts();

	   // Index.initDashboardDaterange();

	   // Index.initIntro();

		// FormComponents.init();

		$('.date-picker').datepicker({
			altFormat: "yy-mm-dd",
			appendText: "(yyyy-mm-dd)",
			dateFormat: "yy-mm-dd",
			format: "yyyy-mm-dd",
			clearBtn: true,
			todayHighlight: true,
			autoclose: true,
			todayBtn: true,
			language: "zh-CN" //汉化
		});

		$('.year-picker').datepicker({
			startView: 2,
            maxViewMode: 2,
            minViewMode:2,
            format: "yyyy",//选择日期后，文本框显示的日期格式
			autoclose: true,
			clearBtn: true,
            language: "zh-CN" //汉化
		});

		$('.month-picker').datepicker({
			format: "yyyy-mm",
			todayBtn: true,
			// startDate: "2013-02-14 10:00:00",
			todayHighlight: true,
			autoclose: true,
			clearBtn: true,
			language: "zh-CN" //汉化
			// pickerPosition: (App.isRTL() ? "bottom-right" : "bottom-left"),
			// minuteStep: 10
		})

		$('.form_datetime').datetimepicker({
			format: "yyyy-mm-dd hh:ii:00",
			todayBtn: true,
			// startDate: "2013-02-14 10:00:00",
			todayHighlight: true,
			autoclose: true,
			clearBtn: true,
			language: "zh-CN" //汉化
			// pickerPosition: (App.isRTL() ? "bottom-right" : "bottom-left"),
			// minuteStep: 10
		})


        checkNum();

        // console.log(DateDiff("2017-09-15", "2017-12-22"))
	});


	// 模态框消失的时候,把模态框的表单重置为初始状态
	$('.modal').on('hidden', function (e) {
		var target = e.currentTarget;
		$(target).find("div.checker>span").removeClass("checked");
		$(target).find("div.checker>span>input[type=checkbox]").attr("checked", false);
		$(target).find('form')[0].reset();
	});

	function pageReload() {
		location.reload();
	}

	// 信息提示框
	function customShowMsg(msg) {
		layer.msg(msg);

		// layer.alert(msg, {
		// 	// icon: 5,
		// 	// skin: 'layer-ext-moon'
		// })
	}

	/**
	 * 自定义二次确认
	 * @param  {string} url
	 * @return none
	 */
	function customConfirm(url) {
		layer.confirm('确认执行此操作吗?', {
		  	btn: ['确认','取消'] //按钮
		}, function(){
			$.ajax({
			    type:'get',
			    dataType: 'json',
			    url: url,
			    success: function (res) {
			        if(res.status == 1) {
			            // parent.closepop(0);
			            // parent.window.location.reload();
			        	if(res.is_parent_reload == 1) {
			        		parent.pageReload();
			        	}

			            pageReload();
			        }else {
			        	customShowMsg(res.message);
			        }
			    },
			    error: function () {
			        customShowMsg("出现未知错误,请联系管理员");
			    }
			});
		});
	}
	/**
	 * 自定义表单二次确认提交
	 * @param  {string} url
	 * @return none
	 */
	function formConfirm(url) {
		layer.confirm('确认执行此操作吗?', {
			btn: ['确认','取消'] //按钮
		}, function(){
			$.ajax({
				type:'get',
				dataType: 'json',
				url: url,
				success: function (res) {
					if(res.status == 1) {
						// parent.closepop(0);
						// parent.window.location.reload();
						if(res.is_parent_reload == 1) {
							parent.pageReload();
						}

						pageReload();
					}else {
						customShowMsg(res.message);
					}
				},
				error: function () {
					customShowMsg("出现未知错误,请联系管理员");
				}
			});
		});
	}

	/**
	 * 导出确认
	 * @param  {string} url
	 * @return none
	 */
	function exportConfirm(url) {
		layer.confirm('确认执行此操作吗?', {
			btn: ['确认','取消'] //按钮
		}, function(index){
			var params = $('#financeListFrom').serialize();

			var start_time = $('#financeListFrom input[name=start_time]').val();
			var end_time = $('#financeListFrom input[name=end_time]').val();

			var dateDiff = DateDiff(end_time, start_time);
			// console.log(dateDiff);return false;


			if((!start_time || !end_time) || dateDiff > 7) {
				customShowMsg("开始时间或者结束时间必填, 并且不能超过7天");
				return false;
			}

			// location.href = ;
			window.open(url+"?"+params);
			layer.close(index);
		});
	}
	/**
	 * 导出确认(用户列表)
	 * @param  {string} url
	 * @return none
	 */
	function userExportConfirm(url) {
		layer.confirm('确认执行此操作吗?', {
			btn: ['确认','取消'] //按钮
		}, function(index){
			var params = $('#financeListFrom').serialize();

			var start_time = $('#financeListFrom input[name=start_time]').val();
			var end_time = $('#financeListFrom input[name=end_time]').val();

			var dateDiff = DateDiff(end_time, start_time);
			// console.log(dateDiff);return false;


//			if((!start_time || !end_time) || dateDiff > 7) {
//				customShowMsg("开始时间或者结束时间必填, 并且不能超过7天");
//				return false;
//			}

			// location.href = ;
			window.open(url+"?"+params);
			layer.close(index);
		});
	}

	//计算天数差的函数，通用
	 function  DateDiff(sDate1,  sDate2){    //sDate1和sDate2是2006-12-18格式
	     var  aDate,  oDate1,  oDate2,  iDays
	     aDate  =  sDate1.split("-")
	     oDate1  =  new  Date(aDate[1]  +  '-'  +  aDate[2]  +  '-'  +  aDate[0])    //转换为12-18-2006格式
	     aDate  =  sDate2.split("-")
	     oDate2  =  new  Date(aDate[1]  +  '-'  +  aDate[2]  +  '-'  +  aDate[0])
	     iDays  =  parseInt(Math.abs(oDate1  -  oDate2)  /  1000  /  60  /  60  /24)    //把相差的毫秒数转换为天数
	     return  iDays
	 }


	function customPrompt(title, url, param) {
		// var data = "";

		layer.prompt({title: title, formType: 2}, function(val, index){
		  	$.ajax({
			    type:'post',
			    dataType: 'json',
			    url: url,
			    data: param+"&remark="+val,
			    success: function (res) {
			        if(res.status == 1) {
			            parent.pageReload();
			        }else {
			        	customShowMsg(res.message);
			        }
			    },
			    error: function () {
			        customShowMsg("出现未知错误,请联系管理员");
			    }
			});

		  	layer.close(index);
		});
	}

	function customPromptAndGetRemark() {
		var remark;
		remark = layer.prompt(function(val, index){
			remark = val;
		  	layer.msg('得到了'+val);
		  	layer.close(index);

		  	return val;
		});

		console.log(remark);

		return remark;
	}
	function hiddenSee(){
		$('#div-page-sidebar').toggle();
	}
	/*
	js判断设备是手机端还是电脑端
	 */
	function browserRedirect() {

		var sUserAgent = navigator.userAgent.toLowerCase();
		var bIsIpad = sUserAgent.match(/ipad/i) == "ipad";
		var bIsIphoneOs = sUserAgent.match(/iphone os/i) == "iphone os";
		var bIsMidp = sUserAgent.match(/midp/i) == "midp";
		var bIsUc7 = sUserAgent.match(/rv:1.2.3.4/i) == "rv:1.2.3.4";
		var bIsUc = sUserAgent.match(/ucweb/i) == "ucweb";
		var bIsAndroid = sUserAgent.match(/android/i) == "android";
		var bIsCE = sUserAgent.match(/windows ce/i) == "windows ce";
		var bIsWM = sUserAgent.match(/windows mobile/i) == "windows mobile";
		//document.writeln("您的浏览设备为：");
		if (bIsIpad || bIsIphoneOs || bIsMidp || bIsUc7 || bIsUc || bIsAndroid || bIsCE || bIsWM) {
			//document.writeln("phone");
			return flag = true;
		} else {
			//document.writeln("pc");
			return flag = false;
		}
	}
	/*
	切换直营店
	 */
	function chooseStore(){


	}

	function checkNum() {
		var ulr = "/recommend/match/ajaxchecknum.html"
		$.ajax({
		    type:'get',
		    dataType: 'json',
		    url: ulr,
		    success: function (res) {
		        if(res.length == 0) {
		        	return false;
		        }
		        var len = res.length;
		        var li1 = "<li><p>"+len+" 条推荐待审核</p></li><li>";
		        var liLast = '<li class="external"><a href="/recommend/audit/auditList.html">查看全部 <i class="m-icon-swapright"></i></a></li>';

		        var li = "";
		        // console.log(res);return;
		        $.each(res, function (i, n) {
		        	var expertName = n.expertName;
		        	// console.log(typeof expertName);
		        	if(expertName == null) {
		        		expertName = "海推";
		        	}
		        	li += '<li>';
					// li += '<a href="javascript:" onclick="showpop(\'审核推荐\', \'/recommend/audit/auditpreview.html?article_id='+n.id+'\')">';
					li += '<a href="javascript:" onclick="showpop(\'审核推荐\', \''+n.audit_url+'\')">';
					li += '<span class="photo">';
					li += 	'<img src="'+n.exp_head+'" alt="">';
					li += '</span>';
					li += '<span class="subject"><span class="from">'+expertName+'</span>';
					// li += '<span class="time">1030小时</span>';
					li += '</span><span class="message">'+n.title+'</span>';
					li += '</a></li>';
		        });

		        $('#hlader_inbox_bar .dropdown-toggle .badge').text(len);

		        $('#hlader_inbox_bar .dropdown-menu').empty();
		        $('#hlader_inbox_bar .dropdown-menu').append(li1);
		        $('#hlader_inbox_bar .dropdown-menu').append(li);
		        $('#hlader_inbox_bar .dropdown-menu').append(liLast);
		    },
		    error: function () {
		        // customShowMsg("出现未知错误,请联系管理员");
		    }
		});
        setTimeout("checkNum()", 10000);
	}

</script>

<script>
 /*
  *flv视频播放
  */
//	if(flvjs.isSupported()){
//		var videoElement = document.getElementById('videoElement');
//		var myPlayer = document.getElementById('my-video');
//		var flvPlayer = flvjs.createPlayer({
//			type:'flv',
//			url:myPlayer.getAttribute("data-url")
//		});
//		flvPlayer.attachMediaElement(videoElement);
//		flvPlayer.load();
////            flvPlayer.play();
//	}
//
//	videoElement.onclick=function(){
//		if(videoElement.paused){
//			videoElement.play();
//		}else{
//			videoElement.pause();
//		}
//	}

// var myPlayer = document.getElementById('my-video');
 var videoElement = document.getElementById('videoElement');
 var snapshot = document.getElementById('my-video').getAttribute("data-snapshot");
 var myPlayer=neplayer("my-video", {
	 poster:snapshot,
	 preload: "auto"
 }, function(){

 }).setDataSource(

	 document.getElementById('videoElement').getAttribute("data-url")
 );

</script>

<!-- END JAVASCRIPTS -->


<script>

	// jQuery(document).ready(function() {

	//    App.init();

	//    TableManaged.init();

	// });

</script>

	<div id="think_page_trace" style="position: fixed;bottom:0;right:0;font-size:14px;width:100%;z-index: 999999;color: #000;text-align:left;font-family:&#39;微软雅黑&#39;;">
    <div id="think_page_trace_tab" style="display: none;background:white;margin:0;height: 250px;">
        <div id="think_page_trace_tab_tit" style="height:30px;padding: 6px 12px 0;border-bottom:1px solid #ececec;border-top:1px solid #ececec;font-size:16px">
                        <span style="color: rgb(0, 0, 0); padding-right: 12px; height: 30px; line-height: 30px; display: inline-block; margin-right: 3px; cursor: pointer; font-weight: 700;">基本</span>
                        <span style="color: rgb(153, 153, 153); padding-right: 12px; height: 30px; line-height: 30px; display: inline-block; margin-right: 3px; cursor: pointer; font-weight: 700;">文件</span>
                        <span style="color: rgb(153, 153, 153); padding-right: 12px; height: 30px; line-height: 30px; display: inline-block; margin-right: 3px; cursor: pointer; font-weight: 700;">流程</span>
                        <span style="color: rgb(153, 153, 153); padding-right: 12px; height: 30px; line-height: 30px; display: inline-block; margin-right: 3px; cursor: pointer; font-weight: 700;">错误</span>
                        <span style="color: rgb(153, 153, 153); padding-right: 12px; height: 30px; line-height: 30px; display: inline-block; margin-right: 3px; cursor: pointer; font-weight: 700;">SQL</span>
                        <span style="color: rgb(153, 153, 153); padding-right: 12px; height: 30px; line-height: 30px; display: inline-block; margin-right: 3px; cursor: pointer; font-weight: 700;">调试</span>
                    </div>
        <div id="think_page_trace_tab_cont" style="overflow:auto;height:212px;padding:0;line-height: 24px">
                        <div style="display: block;">
                <ol style="padding: 0; margin:0">
                    <li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">请求信息 : 2018-05-23 19:26:28 HTTP/1.1 GET : dev.cms.catjc.com/index/index/index</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">运行时间 : 0.080112s [ 吞吐率：12.48req/s ] 内存消耗：4,658.71kb 文件加载：61</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">查询信息 : 3 queries 0 writes </li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">缓存信息 : 3 reads,0 writes</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">配置加载 : 75</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">会话信息 : SESSION_ID=93p0m7org9rakbbunq5o8evoe6</li>                </ol>
            </div>
                        <div style="display:none;">
                <ol style="padding: 0; margin:0">
                    <li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/public/index.php ( 0.81 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/start.php ( 0.71 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/base.php ( 2.60 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/Loader.php ( 18.40 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/Error.php ( 3.47 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/Config.php ( 5.86 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/convention.php ( 9.57 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/App.php ( 20.00 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/Request.php ( 46.45 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/config/config.php ( 10.60 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/config/extra/cacherefresh.php ( 0.63 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/config/extra/database.php ( 2.69 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/config/extra/platform.php ( 0.44 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/config/extra/qiniu.php ( 0.77 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/config/extra/umeng.php ( 1.73 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/application/common.php ( 0.97 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/extend/common/model.php ( 5.34 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/extend/common/time.php ( 5.16 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/extend/common/sys.php ( 6.20 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/extend/common/file.php ( 1.48 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/extend/common/arr.php ( 2.40 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/extend/common/enum.php ( 28.73 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/extend/common/dataformat.php ( 0.75 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/extend/RedisDB.php ( 10.57 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/extend/HMongoDB.php ( 1.90 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/Env.php ( 1.05 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/helper.php ( 16.25 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/Hook.php ( 4.70 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/Lang.php ( 6.77 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/Log.php ( 5.51 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/lang/zh-cn.php ( 3.67 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/config/route.php ( 0.72 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/Route.php ( 57.56 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/application/index/controller/IndexController.php ( 0.22 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/application/BaseController.php ( 13.40 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/Controller.php ( 6.01 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/traits/controller/Jump.php ( 4.75 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/View.php ( 6.67 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/view/driver/Think.php ( 5.45 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/Template.php ( 45.38 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/template/driver/File.php ( 2.17 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/Session.php ( 10.56 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/application/sys/model/ModuleModel.php ( 5.83 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/Model.php ( 48.35 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/Db.php ( 6.36 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/db/connector/Dblib.php ( 4.10 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/db/Connection.php ( 25.74 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/db/Query.php ( 73.15 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/db/builder/Dblib.php ( 3.36 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/db/Builder.php ( 27.51 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/Debug.php ( 6.88 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/exception/ErrorException.php ( 1.83 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/Exception.php ( 1.57 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/exception/Handle.php ( 8.01 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/Cache.php ( 6.13 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/cache/driver/File.php ( 6.71 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/cache/Driver.php ( 5.25 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php ( 24.96 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/Url.php ( 12.24 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/Response.php ( 8.18 KB )</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">/home/wwwroot/dev.cms.catjc.com/thinkphp/library/think/debug/Html.php ( 4.16 KB )</li>                </ol>
            </div>
                        <div style="display:none;">
                <ol style="padding: 0; margin:0">
                    <li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[ LANG ] /home/wwwroot/dev.cms.catjc.com/thinkphp/lang/zh-cn.php</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[ ROUTE ] array (
  'type' =&gt; 'module',
  'module' =&gt;
  array (
    0 =&gt; 'index',
    1 =&gt; 'index',
    2 =&gt; 'index',
  ),
)</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[ HEADER ] array (
  'host' =&gt; 'dev.cms.catjc.com',
  'connection' =&gt; 'keep-alive',
  'cache-control' =&gt; 'max-age=0',
  'upgrade-insecure-requests' =&gt; '1',
  'user-agent' =&gt; 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36',
  'accept' =&gt; 'text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8',
  'referer' =&gt; 'http://dev.cms.catjc.com/index/index/index',
  'accept-encoding' =&gt; 'gzip, deflate',
  'accept-language' =&gt; 'zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7',
  'cookie' =&gt; 'thinkphp_show_page_trace=0|0; Hm_lvt_ed542f707266f465c98441ea08210f68=1525227776,1525682178,1526260288,1526780081; PHPSESSID=93p0m7org9rakbbunq5o8evoe6; Hm_lpvt_ed542f707266f465c98441ea08210f68=1527074714',
  'content-type' =&gt; '',
  'content-length' =&gt; '',
)</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[ PARAM ] array (
)</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[ SESSION ] INIT array (
  'id' =&gt; '',
  'var_session_id' =&gt; '',
  'prefix' =&gt; 'think',
  'type' =&gt; '',
  'auto_start' =&gt; true,
)</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[ DB ] INIT dblib</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[ CACHE ] INIT File</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[ RUN ] app\index\controller\IndexController-&gt;index[ /home/wwwroot/dev.cms.catjc.com/application/index/controller/IndexController.php ]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[ VIEW ] /home/wwwroot/dev.cms.catjc.com/public/../application//view/index/index.html [ array (
  0 =&gt; 'adminInfo',
  1 =&gt; 'adminPriList',
  2 =&gt; 'adminPriListLevel1',
  3 =&gt; 'adminPriListLevel2',
  4 =&gt; 'is_index',
) ]</li>                </ol>
            </div>
                        <div style="display:none;">
                <ol style="padding: 0; margin:0">
                    <li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组下标: 0[/home/wwwroot/dev.cms.catjc.com/application/sys/model/ModuleModel.php:211]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:211]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:211]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:211]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:211]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:211]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:211]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:211]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:211]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:211]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:211]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:211]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:211]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:211]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:211]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:211]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:211]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:211]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:211]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:211]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:211]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:211]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:211]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[8]未定义数组索引: active[/home/wwwroot/dev.cms.catjc.com/runtime/temp/4abcc7f868a41939f74211a254bccbe4.php:216]</li>                </ol>
            </div>
                        <div style="display:none;">
                <ol style="padding: 0; margin:0">
                    <li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[ DB ] CONNECT:[ UseTime:0.002129s ] dblib:host=192.168.1.91:1433;dbname=Catjc;charset=utf8</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[ SQL ] SELECT   column_name,   data_type,   column_default,   is_nullable
        FROM    information_schema.tables AS t
        JOIN    information_schema.columns AS c
        ON  t.table_catalog = c.table_catalog
        AND t.table_schema  = c.table_schema
        AND t.table_name    = c.table_name
        WHERE   t.table_name = 'sys_module' [ RunTime:0.000983s ]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[ SQL ] SELECT column_name FROM information_schema.key_column_usage WHERE table_name='sys_module' [ RunTime:0.000619s ]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[ SQL ] SELECT T1.* FROM (SELECT thinkphp.*, ROW_NUMBER() OVER ( ORDER BY rand()) AS ROW_NUMBER FROM (SELECT  * FROM [sys_module] WHERE  [path] = '/index/Index/index') AS thinkphp) AS T1 WHERE (T1.ROW_NUMBER BETWEEN 1 AND 1) [ RunTime:0.000701s ]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[ SQL ] SELECT * FROM sys_module WHERE id IN (SELECT parent_id FROM sys_module WHERE path = '/index/index/index' AND ([level] = 3 OR [level] = 2)) [ RunTime:0.000460s ]</li><li style="border-bottom:1px solid #EEE;font-size:14px;padding:0 12px">[ SQL ] SELECT T1.* FROM (SELECT thinkphp.*, ROW_NUMBER() OVER ( ORDER BY rand()) AS ROW_NUMBER FROM (SELECT  * FROM [sys_module] WHERE  [is_delete] = 0) AS thinkphp) AS T1 [ RunTime:0.000550s ]</li>                </ol>
            </div>
                        <div style="display:none;">
                <ol style="padding: 0; margin:0">
                                    </ol>
            </div>
                    </div>
    </div>
    <div id="think_page_trace_close" style="display:none;text-align:right;height:15px;position:absolute;top:10px;right:12px;cursor:pointer;"><img style="vertical-align:top;" src="data:image/gif;base64,R0lGODlhDwAPAJEAAAAAAAMDA////wAAACH/C1hNUCBEYXRhWE1QPD94cGFja2V0IGJlZ2luPSLvu78iIGlkPSJXNU0wTXBDZWhpSHpyZVN6TlRjemtjOWQiPz4gPHg6eG1wbWV0YSB4bWxuczp4PSJhZG9iZTpuczptZXRhLyIgeDp4bXB0az0iQWRvYmUgWE1QIENvcmUgNS4wLWMwNjAgNjEuMTM0Nzc3LCAyMDEwLzAyLzEyLTE3OjMyOjAwICAgICAgICAiPiA8cmRmOlJERiB4bWxuczpyZGY9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkvMDIvMjItcmRmLXN5bnRheC1ucyMiPiA8cmRmOkRlc2NyaXB0aW9uIHJkZjphYm91dD0iIiB4bWxuczp4bXA9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC8iIHhtbG5zOnhtcE1NPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvbW0vIiB4bWxuczpzdFJlZj0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL3NUeXBlL1Jlc291cmNlUmVmIyIgeG1wOkNyZWF0b3JUb29sPSJBZG9iZSBQaG90b3Nob3AgQ1M1IFdpbmRvd3MiIHhtcE1NOkluc3RhbmNlSUQ9InhtcC5paWQ6MUQxMjc1MUJCQUJDMTFFMTk0OUVGRjc3QzU4RURFNkEiIHhtcE1NOkRvY3VtZW50SUQ9InhtcC5kaWQ6MUQxMjc1MUNCQUJDMTFFMTk0OUVGRjc3QzU4RURFNkEiPiA8eG1wTU06RGVyaXZlZEZyb20gc3RSZWY6aW5zdGFuY2VJRD0ieG1wLmlpZDoxRDEyNzUxOUJBQkMxMUUxOTQ5RUZGNzdDNThFREU2QSIgc3RSZWY6ZG9jdW1lbnRJRD0ieG1wLmRpZDoxRDEyNzUxQUJBQkMxMUUxOTQ5RUZGNzdDNThFREU2QSIvPiA8L3JkZjpEZXNjcmlwdGlvbj4gPC9yZGY6UkRGPiA8L3g6eG1wbWV0YT4gPD94cGFja2V0IGVuZD0iciI/PgH//v38+/r5+Pf29fTz8vHw7+7t7Ovq6ejn5uXk4+Lh4N/e3dzb2tnY19bV1NPS0dDPzs3My8rJyMfGxcTDwsHAv769vLu6ubi3trW0s7KxsK+urayrqqmop6alpKOioaCfnp2cm5qZmJeWlZSTkpGQj46NjIuKiYiHhoWEg4KBgH9+fXx7enl4d3Z1dHNycXBvbm1sa2ppaGdmZWRjYmFgX15dXFtaWVhXVlVUU1JRUE9OTUxLSklIR0ZFRENCQUA/Pj08Ozo5ODc2NTQzMjEwLy4tLCsqKSgnJiUkIyIhIB8eHRwbGhkYFxYVFBMSERAPDg0MCwoJCAcGBQQDAgEAACH5BAAAAAAALAAAAAAPAA8AAAIdjI6JZqotoJPR1fnsgRR3C2jZl3Ai9aWZZooV+RQAOw=="></div>
</div>
<div id="think_page_trace_open" style="height:30px;float:right;text-align:right;overflow:hidden;position:fixed;bottom:0;right:0;color:#000;line-height:30px;cursor:pointer;">
    <div style="background:#232323;color:#FFF;padding:0 6px;float:right;line-height:30px;font-size:14px">0.081790s </div>
    <img width="30" style="" title="ShowPageTrace" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADAAAAAwCAYAAABXAvmHAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAyBpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuMC1jMDYwIDYxLjEzNDc3NywgMjAxMC8wMi8xMi0xNzozMjowMCAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIiB4bWxuczp4bXBNTT0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL21tLyIgeG1sbnM6c3RSZWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZVJlZiMiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIENTNSBXaW5kb3dzIiB4bXBNTTpJbnN0YW5jZUlEPSJ4bXAuaWlkOjVERDVENkZGQjkyNDExRTE5REY3RDQ5RTQ2RTRDQUJCIiB4bXBNTTpEb2N1bWVudElEPSJ4bXAuZGlkOjVERDVENzAwQjkyNDExRTE5REY3RDQ5RTQ2RTRDQUJCIj4gPHhtcE1NOkRlcml2ZWRGcm9tIHN0UmVmOmluc3RhbmNlSUQ9InhtcC5paWQ6NURENUQ2RkRCOTI0MTFFMTlERjdENDlFNDZFNENBQkIiIHN0UmVmOmRvY3VtZW50SUQ9InhtcC5kaWQ6NURENUQ2RkVCOTI0MTFFMTlERjdENDlFNDZFNENBQkIiLz4gPC9yZGY6RGVzY3JpcHRpb24+IDwvcmRmOlJERj4gPC94OnhtcG1ldGE+IDw/eHBhY2tldCBlbmQ9InIiPz5fx6IRAAAMCElEQVR42sxae3BU1Rk/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">
</div>

<script type="text/javascript">
    (function(){
        var tab_tit  = document.getElementById('think_page_trace_tab_tit').getElementsByTagName('span');
        var tab_cont = document.getElementById('think_page_trace_tab_cont').getElementsByTagName('div');
        var open     = document.getElementById('think_page_trace_open');
        var close    = document.getElementById('think_page_trace_close').children[0];
        var trace    = document.getElementById('think_page_trace_tab');
        var cookie   = document.cookie.match(/thinkphp_show_page_trace=(\d\|\d)/);
        var history  = (cookie && typeof cookie[1] != 'undefined' && cookie[1].split('|')) || [0,0];
        open.onclick = function(){
            trace.style.display = 'block';
            this.style.display = 'none';
            close.parentNode.style.display = 'block';
            history[0] = 1;
            document.cookie = 'thinkphp_show_page_trace='+history.join('|')
        }
        close.onclick = function(){
            trace.style.display = 'none';
            this.parentNode.style.display = 'none';
            open.style.display = 'block';
            history[0] = 0;
            document.cookie = 'thinkphp_show_page_trace='+history.join('|')
        }
        for(var i = 0; i < tab_tit.length; i++){
            tab_tit[i].onclick = (function(i){
                return function(){
                    for(var j = 0; j < tab_cont.length; j++){
                        tab_cont[j].style.display = 'none';
                        tab_tit[j].style.color = '#999';
                    }
                    tab_cont[i].style.display = 'block';
                    tab_tit[i].style.color = '#000';
                    history[1] = i;
                    document.cookie = 'thinkphp_show_page_trace='+history.join('|')
                }
            })(i)
        }
        parseInt(history[0]) && open.click();
        tab_tit[history[1]].click();
    })();
</script>

</body></html>
`