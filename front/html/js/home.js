$(document).ready(function(){
	var user = getUrlParam('name')
	$("#accountName").text(user)
	query(user)
	
	$("#topup").click(function(){
		$("#txType").val("topup");
		$("#amount").attr("placeholder","100 USD");
		$("#fee").val("0 USD");
	});

	$("#invest").click(function(){
		
		$("#txType").val("invest");
		$("#amount").attr("placeholder","100 GP Coins");
		$("#fee").val("5% USD");
	});

	$("#cashout").click(function(){
		$("#txType").val("cashout");
		$("#amount").attr("placeholder","100 GP Coins");
		$("#fee").val("5% USD");
	});
	
	$("#withdrawLimit").click(function(){
		$("#txType").val("withdraw limit");
		$("#amount").attr("placeholder","100 GP Coins");
		$("#fee").val("0 USD");
	});

	$("#transfer").click(function(){
		$(".transfer").show()
		$("#txType2").val("transfer");
		$("#amount").attr("placeholder","100 GP Coins");
		$("#fee").val("0 USD");	
	});
	
	$("#middleTx").click(function(){
		  user = getUrlParam("name")
			func   = $("#txType").val()
			amount = $("#amount").val()
			args = {
				"User" : user,
				"Amount": amount,
				}
			middleTx(func, args)
			
			 $('.dialog').hide();
	});
	
	$("#transferTx").click(function(){
			func = $("#txType2").val()
			user = getUrlParam("name")
			amount   = $("#amount2").val()
			receiver = $("#party").val()
			args = {
			"From"   :   user,
			"Amount" :   amount,
			"To" : receiver,	
			}
			transferTx(func,args)
			$(".dialog").hide();
			$(".dialog2").hide();
	});

	$("#blcokInfo").click(function(){
			window.open("./page2.html");		
	});
	
	$("#detial").click(function(){
			query(user);
	});	
	
	setInterval(function(){query(user)}, 3000);	
})

function query(user){
		$.post("http://47.92.154.23:8080/query", {"User": user},
			function(result){
					$("#GPCoin").text(result.gpcoin);
					$("#USD").text(result.usd);
			}
		);
}
/*	
		if(user == "Luke"){
			$.get("http://47.92.154.23:8080/loan", function(result){
					$("#loan").text(result.amount)
			
			});
		}else{
			$.get("http://47.92.154.23:8080/limits", function(result){
					$("#limit").text(result.amount)
			
			});

		}

	
}

*/
function middleTx(func, args){
	$.post("http://47.92.154.23:8080/"+func, args);
}
function transferTx(func, args){
	if(func == "transfer limit"){
		$.post("http://47.92.154.23:8080/tLimits", args);
	}else{
		$.post("http://47.92.154.23:8080/transfer", args);
	}//alert(args.From+args.Amount+args.Receiver)
}


function getUrlParam(name) {
            var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); 
            var r = window.location.search.substr(1).match(reg);  
            if (r != null) return unescape(r[2]); return null; 
        }
