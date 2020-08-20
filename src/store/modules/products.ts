import { VuexModule, Module } from "vuex-module-decorators";
import Product from "@/models/Product";
import ProductBrand from '@/models/ProductBrand';
import ProductType from '@/models/ProductType';
@Module({
    namespaced: true,
    name: "Products"
})
export default class Products extends VuexModule {
    private productsList: Array<Product> = [];
    private productsBrandList: Array<ProductBrand> = [];
    private productsTypeList: Array<ProductType> = [];
}
