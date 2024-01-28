<template>
  <div class="mt-3">
    <div class="row">
        <!-- Table -->
        <div class="col-xl-12">
            <Portlet :headertitle="tableTitle">
                <div class="card-body pt-0">
                    <div class="table-responsive mb-0">
                        <table class="table table-hover table-centered mb-0">
                            <thead>
                                <tr>
                                    <th>Product Id</th>
                                    <th>Product Name</th>
                                    <th>Price</th>
                                    <th>Quantity</th>
                                    <th>Amount</th>
                                    <th>Date</th>
                                    <th>Sales</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr v-for="sellingData in productData" :key="sellingData.id">
                                    <td>{{ sellingData.productid }}</td>
                                    <td>{{ sellingData.name }}</td>
                                    <td>{{ sellingData.price }}</td>
                                    <td>{{ sellingData.quantity }}</td>
                                    <td>{{ sellingData.amount }}</td>
                                    <td>{{ sellingData.date }}</td>
                                    <td>{{ sellingData.sales }}</td>
                                </tr>
                            </tbody>
                        </table>
                    </div>
                </div>
            </Portlet>
        </div>
    </div>
  </div>
</template>

<script>
export default {
  data(){
    return {
      tableTitle: "Top Selling Products",
            productData: [{
                    name: "ASOS Ridley High Waist",
                    price: "$79.49",
                    quantity: 82,
                    amount: "$6,518.18",
                    date: "Sep 1, 2018",
                    sales: 54,
                    productid: 200125
                },
                {
                    name: "Marco Lightweight Shirt",
                    price: "$128.50",
                    quantity: 37,
                    amount: "$4,754.50",
                    date: "Sep 15, 2018",
                    sales: 28,
                    productid: 200130
                },
                {
                    name: "Half Sleeve Shirt",
                    price: "$39.99",
                    quantity: 64,
                    amount: "$2,559.36",
                    date: "Nov 1, 2018",
                    sales: 55,
                    productid: 200140
                },
                {
                    name: "Lightweight Jacket",
                    price: "$20.00",
                    quantity: 184,
                    amount: "$3,680.00",
                    date: "Nov 15, 2018",
                    sales: 85,
                    productid: 200250
                },
                {
                    name: "Marco Shoes",
                    price: "$28.49",
                    quantity: 69,
                    amount: "$1,965.81",
                    date: "Jan 1, 2019",
                    sales: 49,
                    productid: 200256
                },
                {
                    name: "ASOS Ridley High Waist",
                    price: "$79.49",
                    quantity: 82,
                    amount: "$6,518.18",
                    date: "Sep 1, 2018",
                    sales: 54,
                    productid: 200125
                },
                {
                    name: "Half Sleeve Shirt",
                    price: "$39.99",
                    quantity: 64,
                    amount: "$2,559.36",
                    date: "Nov 1, 2018",
                    sales: 55,
                    productid: 200140
                },
                {
                    name: "Lightweight Jacket",
                    price: "$20.00",
                    quantity: 184,
                    amount: "$3,680.00",
                    date: "Nov 15, 2018",
                    sales: 85,
                    productid: 200250
                }
            ],
    }
  }
}
</script>
