package main

import (
	"fmt"
	"strconv"
	"strings"
)

var firstInput = "183\n185\n186\n182\n185\n191\n204\n203\n205\n220\n215\n208\n216\n217\n218\n219\n210\n226\n227\n231\n245\n243\n255\n257\n250\n255\n286\n287\n286\n293\n294\n296\n299\n316\n322\n330\n332\n337\n340\n341\n339\n348\n349\n350\n351\n354\n357\n359\n360\n361\n363\n364\n365\n372\n373\n371\n374\n375\n379\n383\n384\n383\n388\n390\n395\n397\n410\n412\n411\n414\n411\n432\n426\n429\n437\n438\n462\n463\n464\n477\n480\n497\n516\n524\n536\n547\n550\n555\n557\n556\n565\n566\n567\n574\n575\n583\n585\n594\n596\n617\n621\n616\n613\n616\n618\n614\n615\n618\n619\n621\n616\n644\n649\n650\n655\n658\n657\n658\n661\n668\n647\n651\n657\n654\n656\n657\n658\n666\n676\n686\n692\n701\n706\n705\n706\n697\n701\n670\n663\n667\n668\n676\n677\n676\n692\n714\n715\n734\n768\n785\n824\n829\n845\n844\n847\n860\n865\n880\n893\n895\n899\n900\n902\n905\n906\n907\n906\n902\n912\n927\n928\n945\n952\n953\n961\n969\n970\n974\n975\n976\n980\n989\n990\n993\n1006\n1008\n1009\n1010\n1012\n1013\n1014\n1017\n1024\n1038\n1058\n1059\n1065\n1066\n1070\n1083\n1085\n1086\n1094\n1097\n1102\n1121\n1122\n1137\n1138\n1145\n1148\n1159\n1170\n1171\n1172\n1179\n1175\n1179\n1178\n1190\n1196\n1215\n1220\n1209\n1213\n1214\n1209\n1211\n1251\n1255\n1269\n1272\n1273\n1284\n1286\n1312\n1314\n1317\n1321\n1341\n1357\n1365\n1366\n1369\n1379\n1359\n1332\n1342\n1362\n1365\n1375\n1378\n1379\n1381\n1387\n1397\n1401\n1400\n1401\n1429\n1453\n1454\n1455\n1461\n1466\n1467\n1463\n1475\n1489\n1488\n1496\n1506\n1515\n1538\n1545\n1538\n1544\n1553\n1555\n1558\n1567\n1573\n1577\n1582\n1583\n1585\n1586\n1585\n1586\n1601\n1613\n1612\n1630\n1632\n1631\n1638\n1642\n1655\n1656\n1613\n1615\n1620\n1611\n1609\n1615\n1617\n1621\n1624\n1634\n1639\n1644\n1655\n1678\n1681\n1683\n1695\n1698\n1708\n1709\n1712\n1713\n1717\n1743\n1744\n1747\n1781\n1782\n1799\n1803\n1804\n1806\n1808\n1825\n1830\n1831\n1833\n1848\n1849\n1851\n1853\n1858\n1857\n1865\n1867\n1886\n1888\n1892\n1902\n1923\n1925\n1930\n1931\n1910\n1916\n1922\n1924\n1925\n1930\n1918\n1925\n1926\n1929\n1934\n1932\n1926\n1932\n1931\n1950\n1952\n1953\n1964\n1965\n1967\n1977\n1979\n1983\n1993\n1994\n1996\n2005\n2020\n2032\n2033\n2039\n2040\n2041\n2045\n2058\n2059\n2061\n2063\n2086\n2089\n2097\n2095\n2101\n2102\n2105\n2106\n2100\n2113\n2158\n2160\n2162\n2164\n2165\n2172\n2174\n2176\n2183\n2184\n2190\n2189\n2212\n2225\n2233\n2192\n2194\n2215\n2222\n2223\n2233\n2241\n2243\n2235\n2246\n2287\n2292\n2297\n2298\n2299\n2300\n2306\n2308\n2318\n2322\n2333\n2332\n2312\n2309\n2313\n2297\n2306\n2323\n2336\n2339\n2322\n2323\n2327\n2332\n2351\n2361\n2363\n2364\n2365\n2371\n2395\n2393\n2394\n2423\n2435\n2438\n2443\n2444\n2450\n2449\n2453\n2456\n2457\n2456\n2462\n2463\n2477\n2478\n2484\n2485\n2498\n2499\n2497\n2498\n2499\n2501\n2502\n2505\n2518\n2521\n2522\n2520\n2521\n2523\n2529\n2530\n2528\n2544\n2545\n2546\n2547\n2590\n2612\n2620\n2631\n2634\n2633\n2634\n2635\n2661\n2663\n2667\n2670\n2676\n2678\n2654\n2664\n2665\n2671\n2681\n2684\n2711\n2725\n2726\n2728\n2729\n2734\n2735\n2740\n2759\n2758\n2762\n2763\n2765\n2768\n2749\n2750\n2752\n2758\n2728\n2732\n2749\n2748\n2763\n2766\n2767\n2769\n2764\n2765\n2768\n2769\n2772\n2758\n2777\n2788\n2811\n2815\n2820\n2824\n2832\n2839\n2841\n2839\n2840\n2869\n2872\n2847\n2857\n2858\n2860\n2865\n2855\n2851\n2866\n2870\n2871\n2872\n2878\n2894\n2899\n2900\n2907\n2913\n2905\n2916\n2917\n2916\n2913\n2919\n2933\n2935\n2940\n2945\n2953\n2965\n2976\n2986\n2990\n3013\n3017\n3004\n2997\n2998\n2999\n3000\n2995\n2996\n2997\n2998\n3006\n3008\n3019\n3022\n3024\n3025\n3026\n3027\n3028\n3031\n3033\n3038\n3042\n3079\n3080\n3087\n3117\n3119\n3120\n3121\n3122\n3123\n3130\n3135\n3130\n3132\n3134\n3153\n3167\n3168\n3173\n3178\n3180\n3189\n3190\n3209\n3242\n3246\n3259\n3265\n3269\n3275\n3279\n3285\n3301\n3318\n3315\n3324\n3328\n3330\n3352\n3354\n3351\n3358\n3392\n3393\n3416\n3417\n3421\n3420\n3441\n3458\n3476\n3481\n3491\n3510\n3512\n3515\n3550\n3551\n3555\n3557\n3586\n3604\n3609\n3616\n3623\n3626\n3606\n3627\n3636\n3637\n3640\n3626\n3635\n3644\n3661\n3665\n3667\n3665\n3671\n3674\n3684\n3688\n3689\n3693\n3720\n3728\n3731\n3730\n3733\n3737\n3751\n3753\n3754\n3755\n3760\n3762\n3765\n3769\n3770\n3769\n3774\n3753\n3755\n3758\n3787\n3801\n3800\n3826\n3840\n3843\n3857\n3858\n3859\n3862\n3901\n3911\n3913\n3915\n3916\n3914\n3896\n3924\n3925\n3927\n3929\n3931\n3935\n3936\n3944\n3945\n3943\n3946\n3948\n3949\n3966\n3967\n3956\n3964\n3971\n3981\n3988\n3989\n3992\n4000\n4007\n4009\n4010\n4018\n4019\n4020\n4021\n4022\n4037\n4025\n4026\n4024\n4025\n4026\n4025\n4028\n4021\n4010\n4038\n4039\n4041\n4042\n4030\n4027\n4000\n4001\n4012\n4013\n4021\n4022\n4023\n4036\n4037\n4051\n4047\n4054\n4057\n4064\n4056\n4040\n4042\n4045\n4060\n4061\n4065\n4066\n4058\n4061\n4058\n4060\n4061\n4065\n4066\n4060\n4062\n4054\n4055\n4056\n4078\n4083\n4084\n4098\n4101\n4102\n4079\n4089\n4106\n4111\n4112\n4147\n4157\n4158\n4165\n4174\n4178\n4183\n4184\n4196\n4197\n4202\n4200\n4202\n4209\n4221\n4225\n4233\n4239\n4240\n4234\n4236\n4249\n4250\n4252\n4253\n4254\n4253\n4254\n4267\n4272\n4281\n4256\n4255\n4256\n4258\n4264\n4265\n4264\n4266\n4270\n4271\n4273\n4284\n4292\n4293\n4292\n4293\n4294\n4295\n4303\n4304\n4309\n4310\n4311\n4313\n4314\n4317\n4322\n4324\n4328\n4322\n4345\n4347\n4354\n4348\n4349\n4352\n4368\n4378\n4385\n4391\n4407\n4416\n4417\n4436\n4439\n4438\n4427\n4436\n4437\n4451\n4452\n4456\n4462\n4464\n4465\n4466\n4475\n4486\n4498\n4499\n4501\n4502\n4511\n4521\n4524\n4525\n4528\n4531\n4535\n4540\n4542\n4544\n4543\n4544\n4558\n4580\n4581\n4582\n4583\n4585\n4596\n4597\n4590\n4597\n4598\n4631\n4649\n4650\n4661\n4662\n4666\n4657\n4661\n4668\n4661\n4663\n4664\n4687\n4697\n4710\n4735\n4742\n4764\n4765\n4774\n4763\n4786\n4787\n4805\n4806\n4808\n4809\n4810\n4838\n4844\n4845\n4877\n4881\n4898\n4899\n4912\n4914\n4915\n4931\n4932\n4946\n4952\n4951\n4952\n4958\n4964\n4972\n4967\n4968\n4966\n4973\n4974\n4979\n4980\n4990\n5008\n5012\n5013\n5019\n5011\n5012\n5013\n5014\n5016\n5032\n5035\n5034\n5035\n5041\n5061\n5062\n5064\n5076\n5094\n5076\n5078\n5080\n5083\n5095\n5094\n5097\n5096\n5106\n5108\n5111\n5115\n5125\n5129\n5136\n5137\n5142\n5156\n5158\n5168\n5170\n5175\n5178\n5191\n5189\n5190\n5198\n5200\n5207\n5206\n5207\n5196\n5226\n5227\n5228\n5235\n5261\n5262\n5266\n5271\n5272\n5273\n5287\n5288\n5294\n5316\n5324\n5325\n5322\n5331\n5329\n5330\n5331\n5335\n5350\n5366\n5372\n5373\n5375\n5377\n5391\n5392\n5393\n5395\n5398\n5407\n5420\n5416\n5420\n5413\n5414\n5420\n5423\n5435\n5439\n5446\n5453\n5478\n5502\n5505\n5507\n5522\n5525\n5527\n5528\n5530\n5543\n5546\n5549\n5553\n5557\n5559\n5560\n5561\n5578\n5560\n5556\n5565\n5566\n5577\n5584\n5586\n5591\n5595\n5596\n5604\n5571\n5578\n5580\n5582\n5583\n5588\n5591\n5590\n5596\n5601\n5604\n5602\n5596\n5628\n5644\n5645\n5647\n5648\n5649\n5655\n5662\n5664\n5672\n5680\n5691\n5702\n5703\n5709\n5712\n5729\n5732\n5747\n5746\n5755\n5777\n5768\n5746\n5747\n5751\n5746\n5751\n5774\n5775\n5782\n5785\n5791\n5798\n5805\n5807\n5808\n5827\n5847\n5856\n5828\n5820\n5807\n5811\n5825\n5826\n5828\n5827\n5851\n5852\n5856\n5855\n5856\n5859\n5857\n5858\n5863\n5879\n5881\n5887\n5907\n5913\n5914\n5919\n5940\n5951\n5964\n5989\n5991\n5993\n5979\n5989\n5990\n6004\n6005\n6004\n6006\n6008\n6009\n6006\n6005\n6006\n5977\n5978\n5979\n5983\n5984\n5987\n5988\n5966\n5965\n5970\n5961\n5968\n5969\n5982\n5997\n6006\n6007\n6021\n6025\n6028\n6030\n6033\n6034\n6035\n6039\n6040\n6053\n6062\n6094\n6097\n6099\n6104\n6109\n6125\n6126\n6137\n6133\n6151\n6174\n6173\n6175\n6180\n6181\n6182\n6184\n6185\n6189\n6190\n6189\n6191\n6192\n6224\n6228\n6229\n6237\n6231\n6236\n6239\n6250\n6264\n6265\n6253\n6262\n6263\n6268\n6269\n6274\n6275\n6276\n6278\n6321\n6322\n6343\n6332\n6348\n6352\n6359\n6368\n6371\n6372\n6390\n6391\n6393\n6398\n6412\n6423\n6424\n6438\n6441\n6445\n6447\n6458\n6473\n6485\n6489\n6517\n6519\n6529\n6531\n6532\n6533\n6543\n6547\n6548\n6547\n6549\n6551\n6553\n6554\n6565\n6566\n6576\n6578\n6584\n6614\n6617\n6620\n6622\n6623\n6625\n6638\n6643\n6660\n6662\n6664\n6637\n6635\n6633\n6638\n6652\n6649\n6666\n6680\n6662\n6681\n6683\n6694\n6695\n6699\n6705\n6707\n6714\n6722\n6728\n6729\n6736\n6746\n6750\n6761\n6763\n6754\n6757\n6759\n6763\n6762\n6763\n6764\n6766\n6790\n6808\n6813\n6815\n6816\n6818\n6819\n6820\n6828\n6826\n6829\n6832\n6854\n6857\n6865\n6861\n6865\n6889\n6890\n6904\n6903\n6905\n6906\n6934\n6935\n6944\n6949\n6951\n6962\n6963\n6972\n6977\n6974\n6973\n6975\n7010\n7009\n7011\n7041\n7044\n7057\n7077\n7094\n7090\n7098\n7096\n7088\n7098\n7115\n7118\n7124\n7146\n7165\n7168\n7176\n7190\n7213\n7201\n7202\n7217\n7222\n7243\n7244\n7248\n7250\n7259\n7255\n7253\n7254\n7260\n7271\n7281\n7293\n7311\n7317\n7323\n7324\n7337\n7346\n7349\n7372\n7367\n7374\n7379\n7389\n7381\n7388\n7389\n7396\n7413\n7415\n7417\n7418\n7428\n7430\n7433\n7434\n7423\n7426\n7427\n7419\n7420\n7421\n7422\n7428\n7429\n7434\n7442\n7445\n7414\n7417\n7419\n7421\n7422\n7425\n7428\n7429\n7435\n7444\n7456\n7459\n7462\n7460\n7475\n7482\n7494\n7493\n7497\n7498\n7513\n7514\n7522\n7524\n7525\n7529\n7531\n7537\n7519\n7521\n7528\n7541\n7542\n7543\n7545\n7546\n7547\n7548\n7550\n7560\n7561\n7563\n7533\n7537\n7539\n7542\n7554\n7556\n7560\n7562\n7566\n7594\n7615\n7626\n7631\n7626\n7627\n7626\n7628\n7636\n7641\n7642\n7643\n7654\n7671\n7675\n7674\n7680\n7681\n7682\n7687\n7693\n7687\n7689\n7694\n7697\n7709\n7730\n7727\n7729\n7733\n7734\n7740\n7750\n7752\n7743\n7750\n7760\n7791\n7801\n7802\n7804\n7806\n7807\n7835\n7834\n7835\n7838\n7841\n7842\n7836\n7837\n7845\n7851\n7863\n7864\n7893\n7894\n7904\n7908\n7909\n7910\n7911\n7910\n7912\n7914\n7918\n7919\n7921\n7936\n7948\n7950\n7952\n7956\n7967\n7973\n7979\n7978\n7979\n7978\n7981\n7982\n7985\n7988\n7989\n7997\n7999\n8007\n8006\n8007\n8016\n8017\n8041\n8020\n8031\n8032\n8033\n8032\n8038\n8039\n8040\n8054\n8055\n8056\n8063\n8062\n8063\n8064\n8055\n8060\n8061\n8068\n8077\n8070\n8074\n8078\n8091\n8105\n8104\n8105\n8104\n8102\n8103\n8112\n8115\n8149\n8147\n8150\n8151\n8170\n8171\n8174\n8175\n8182\n8214\n8189\n8204\n8206\n8213\n8218\n8192\n8182\n8179\n8178\n8181\n8183\n8202\n8209\n8210\n8213\n8217\n8234\n8235\n8239\n8240\n8242\n8223\n8218\n8233\n8236\n8247\n8249\n8252\n8253\n8259\n8276\n8299\n8303\n8309\n8310\n8315\n8321\n8319\n8321\n8323\n8324\n8325\n8335\n8333\n8335\n8334\n8335\n8336\n8352\n8360\n8367\n8370\n8371\n8374\n8386\n8388\n8399\n8402\n8403\n8404\n8405\n8409\n8446\n8449\n8469\n8470\n8471\n8469\n8460\n8463\n8485\n8482\n8470\n8472\n8455\n8454\n8447\n8448\n8450\n8462\n8463\n8476\n8507\n8510\n8518\n8527\n8539\n8550\n8561\n8559\n8562\n8577\n8581\n8607\n8620\n8628\n8629\n8627\n8636\n8644\n8637\n8655\n8656\n8692\n8719\n8720\n8721\n8738\n8741\n8742\n8744\n8747\n8737\n8735\n8736\n8743\n8749\n8758\n8761\n8763\n8794\n8793\n8804\n8807\n8808\n8806\n8831\n8841\n8832\n8835\n8843\n8825\n8826\n8840\n8844\n8845\n8855\n8858\n8879\n8882\n8893\n8895\n8900\n8906\n8917\n8921\n8922\n8931\n8932\n8934\n8937\n8939\n8955\n8949\n8950\n8952\n8953\n8963\n8964\n8989\n8993\n8998\n9002\n8999\n9001\n9002\n9016\n9024\n9029\n9032\n9031\n9039\n9049\n9030\n9033\n9035\n9036\n9044\n9051\n9052\n9057\n9067\n9092\n9101\n9123\n9130\n9124\n9133\n9143\n9149\n9172\n9173\n9178\n9193\n9184\n9189\n9169\n9186\n9187\n9189\n9195\n9197\n9198\n9196\n9218\n9219\n9235\n9267\n9271\n9273\n9278\n9289\n9290\n9307\n9319\n9324\n9328\n9342\n9343\n9344\n9348\n9360\n9366\n9373\n9397\n9392\n9393\n9409\n9411\n9409\n9421\n9430\n9451\n9452\n9453\n9456\n9457\n9466\n9467\n9470\n9496\n9500\n9515\n9518\n9514\n9534\n9541\n9542\n9545\n9551\n9583\n9581\n9573\n9588\n9601\n9591\n9592\n9593\n9600\n9599\n9605\n9610\n9612\n9614\n9627\n9629\n9630\n9636\n9637\n9638\n9639\n9640\n9643\n9635\n9637\n9645\n9649\n9650\n9647\n9648\n9652\n9653\n9663\n9656\n9661\n9687\n9688\n9689\n9690\n9694\n9695\n9702\n9703\n9705\n9711\n9714\n9718\n9742\n9768\n9747\n9757\n9758\n9768\n9771\n9768\n9772\n9779\n9780\n9785\n9804\n9805\n9803\n9801\n9802\n9810\n9816\n9820\n9810\n9815\n9816\n9823\n9824\n9830\n9831\n9841\n9860\n9875"

func main() {
	list := strings.Split(firstInput, "\n")
	previous := int64(0)
	count := 0

	for key, depth := range list {
		depthNum, err := strconv.ParseInt(depth, 10, 64)
		if err != nil {
			fmt.Println("Error Converting string")
			return
		}
		if key > 0 && previous < depthNum {
			count += 1
		}
		previous = depthNum
	}

	// use a sliding window to calculate the sum and count the increases
	intList := []int64{}
	for _, val := range list {
		num, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			fmt.Println("Error Converting string")
			return
		}
		intList = append(intList, num)
	}

	prevSum := int64(0)
	sumCount := 0
	for i := 0; i < len(intList) - 2; i++ {
		first := intList[i]
		second := intList[i+1]
		third := intList[i+2]
		sum := first + second + third

		if i > 0 && sum > prevSum {
			sumCount += 1
		}
		prevSum = sum
	}

	fmt.Println("Result", count)
	fmt.Println("Result Sums", sumCount)
}