package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 0; i < 25; i++ {
		stri := strconv.Itoa(i)
		sql := "CREATE TABLE `xes_level_one_class_stu_"+ stri +"` (\n  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',\n  `es_id` varchar(50) NOT NULL DEFAULT '' COMMENT 'es中的id,用户接流更新数据',\n  `class_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '班级id',\n  `stu_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',\n  `course_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '课程id',\n  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '进班时间',\n  `subject_ids` varchar(50) NOT NULL DEFAULT '' COMMENT '学科id',\n  `is_new` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否是新学员',\n  `interact_right_rate` varchar(10) NOT NULL DEFAULT '' COMMENT '互动题正确率',\n  `interact_join_rate` varchar(10) NOT NULL DEFAULT '' COMMENT '互动题正确率',\n  `month_finish_rate` varchar(10) NOT NULL DEFAULT '' COMMENT '上个自然月到现在发生场次的完课率',\n  `renewal_pro` varchar(25) NOT NULL DEFAULT '' COMMENT '转换意愿类型(short_lower_tag:低,short_middle_tag:中,short_high_tag:高)',\n  `renewal_result` varchar(600) NOT NULL DEFAULT '' COMMENT '转换意愿原因',\n  `return_result` varchar(600) NOT NULL DEFAULT '' COMMENT '退费原因(来源数据部门)',\n  `course_finish_rate` varchar(10) NOT NULL DEFAULT '' COMMENT '课程完课率',\n  `homework_right_rate` varchar(10) NOT NULL DEFAULT '' COMMENT '作业正确率',\n  `homework_finish_num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '作业完成数量',\n  `homework_commit_num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '作业提交数量',\n  `homework_need_num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '需要提交作业数量',\n  `live_need` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '需要出勤总讲数',\n  `live_num` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '出勤总讲数',\n  `unwatch_num` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '未听课讲数',\n  `finish_need` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '完课需要完课讲次',\n  `finish_num` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '完课-完成讲次',\n  `mornread_num` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '晨读完成任务数',\n  `mornread_today` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '今天是否晨读',\n  `mornread_seven_today` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '7天晨读的天数',\n  `mornread_all` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '晨读的总天数',\n  `mornread_avg_score` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '晨读平均分',\n  `modify_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '记录修改时间',\n  PRIMARY KEY (`id`),\n  KEY `class_stu_id` (`class_id`,`stu_id`) USING BTREE COMMENT '班级学员id',\n  KEY `es_id` (`es_id`) USING BTREE\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='班级学员维度核心数据-卢洪光';"
		fmt.Println(sql)
	}
}
