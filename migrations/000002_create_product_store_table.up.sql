CREATE Table product_store(
  product_id int REFERENCES products(id),
  store_id int
);