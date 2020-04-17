package main

func main() {

	var sqlMaps map [string]string

	sqlMaps["a"] ="SELECT functions.name as function_name,park_opened_functions.*,functions.* FROM `park_opened_functions` INNER JOIN functions ON functions.id = park_opened_functions.function_id WHERE (1=1 and park_opened_functions.park_code  in (7100000360)) AND `park_opened_functions`.`function_id` IN (?,?)";
	sqlMaps["b"]="SELECT * FROM t_park_record use index(`key`)  WHERE (vpl_number = ? and `status` = 1) "
	sqlMaps["c"]="INSERT INTO t_pay_success_record SET pay_serial='130100BwEsrKDZac1587006797610551',channel_order='1059102008697271E1587006797611',pay_channel=8,pay_way='1',pay_cid='244',pay_time='1587006806',pay_amount=500,pay_src_amount=500,park_code=3410000005,uid=186968543,ctime=1587006806,total_fee=500,business_type=1,business_id='1026569',scene='0x01';"

	/**

	 */



}
