package com.abfinance.api.domain.trade;

import com.abfinance.api.client.config.ABFinanceApiConfig;
import com.abfinance.api.client.domain.CategoryType;
import com.abfinance.api.client.domain.TradeOrderType;
import com.abfinance.api.client.domain.trade.Side;
import com.abfinance.api.client.domain.trade.TimeInForce;
import com.abfinance.api.client.domain.trade.request.TradeOrderRequest;
import com.abfinance.api.client.restApi.ABFinanceApiTradeRestClient;
import com.abfinance.api.client.service.ABFinanceApiClientFactory;
import org.junit.Test;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.TimeUnit;

public class OpenOrderTest {
    ABFinanceApiTradeRestClient client = ABFinanceApiClientFactory.newInstance("YOUR_API_KEY", "YOUR_API_SECRET", ABFinanceApiConfig.TESTNET_DOMAIN).newTradeRestClient();

    @Test
    public void Test_PlaceSpotOrder() throws InterruptedException {

       List<String> orderLinkIdList = new ArrayList<>();
        var newSpotOrderRequest = TradeOrderRequest.builder().category(CategoryType.SPOT).symbol("XRPUSDT").side(Side.BUY).price("0.1").orderType(TradeOrderType.LIMIT).qty("10").timeInForce(TimeInForce.IOC);

        // place 510 spot orders in canceled status
        int lastIdx = 510;
        for (int i = 0; i < lastIdx; i++) {
            String orderCustomId = "abfinance_test_spot_order_link_id_" + String.valueOf(i + 1);
            var newOrder = client.createOrder(newSpotOrderRequest.orderLinkId(orderCustomId).build());
            System.out.println("order " + (i + 1));
            System.out.println(newOrder);
            orderLinkIdList.add(orderCustomId);
        }

        // after 10 seconds
        TimeUnit.SECONDS.sleep(10);

        var openOrderRequest = TradeOrderRequest.builder();
        //get spot open orders
        var openSpotOrdersResult = client.getOpenOrders(openOrderRequest.category(CategoryType.SPOT).openOnly(1).build());
        System.out.println(openSpotOrdersResult);
        //test spot order
        String lastSpotOrderLinkId = orderLinkIdList.get(lastIdx - 1), firstSpotOrderLinkId = orderLinkIdList.get(0);
        //get 510th spot order
        var lastSpotOrderResult = client.getOpenOrders(openOrderRequest.category(CategoryType.SPOT).orderLinkId(lastSpotOrderLinkId).build());
        System.out.println(lastSpotOrderResult);
        //get 1st spot order
        var firstSpotOrderResult = client.getOpenOrders(openOrderRequest.category(CategoryType.SPOT).orderLinkId(firstSpotOrderLinkId).build());
        System.out.println(firstSpotOrderResult);
    }
}
