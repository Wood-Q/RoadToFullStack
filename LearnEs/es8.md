# ES8 Store 使用文档

## IndexerConfig 配置说明

### 必需字段

1. `Client *elasticsearch.Client`

   - 类型：Elasticsearch 客户端实例
   - 说明：必须提供有效的 Elasticsearch 客户端实例
   - 示例：

   ```go
   client, err := elasticsearch.NewClient(elasticsearch.Config{
       Addresses: []string{"http://localhost:9200"},
   })
   ```

2. `Index string`

   - 类型：字符串
   - 说明：索引名称，不能为空
   - 示例：`"my_index"`

3. `DocumentToFields func(ctx context.Context, doc *schema.Document) (field2Value map[string]FieldValue, err error)`
   - 类型：函数
   - 说明：用于将文档转换为 Elasticsearch 字段的映射函数
   - 必须实现此函数来定义文档如何映射到 ES 字段

### 可选字段

1. `BatchSize int`

   - 类型：整数
   - 默认值：5
   - 说明：控制批量嵌入的最大文本数量
   - 限制：不能超过 1000
   - 示例：`10`

2. `Embedding embedding.Embedder`

   - 类型：嵌入器接口
   - 说明：用于向量化的嵌入方法
   - 在以下两种情况下必须提供：
     - VectorFields 包含除文档内容外的字段
     - VectorFields 包含文档内容且文档的 extra 中未提供向量

3. `LocalMapping *types.TypeMapping`

   - 类型：类型映射
   - 说明：定义索引结构的本地映射
   - 默认值：如果未提供，将使用默认映射（包含 content、extra_location 和 content_dense_vector 字段）

4. `Dynamic *dynamicmapping.DynamicMapping`

   - 类型：动态映射设置
   - 说明：控制索引的动态映射行为
   - 可选值：true、false、strict

5. `ValidationMode MappingValidationMode`

   - 类型：映射验证模式
   - 默认值：ValidationModeError
   - 可选值：
     - ValidationModeError：验证失败时抛出错误
     - ValidationModeWarn：验证失败时记录警告
     - ValidationModeSkip：跳过验证

6. `EnableSchemaCheck bool`
   - 类型：布尔值
   - 默认值：false
   - 说明：是否在索引前启用文档模式验证

## 验证失败情况说明

在使用过程中，可能会遇到以下几种验证失败的情况：

### 1. 字段映射验证失败

- 本地定义的字段在远程索引中不存在
- 字段类型不匹配（例如：期望 text 类型，但远程为其他类型）
- 向量字段维度不匹配
- 向量相似度算法不匹配

### 2. 动态映射设置验证失败

- 本地配置了 dynamic 设置，但远程索引未设置
- 本地和远程的 dynamic 设置不一致

### 3. 文档模式验证失败（当 EnableSchemaCheck 为 true 时）

- 文档缺少必需的字段
- 文档包含未定义的字段
- 字段值类型不符合预期

### 4. 批量处理验证失败

- 批量大小超过限制（>1000）
- 需要嵌入的字段数量超过 BatchSize 限制
- EmbedKey 与现有字段名冲突
- EmbedKey 重复使用

### 5. 向量化相关验证失败

- 需要向量化但未提供 Embedding 实现
- 向量化结果长度与输入文本数量不匹配
- 无法将值转换为字符串进行向量化

## 使用示例

```go
config := &IndexerConfig{
    Client: esClient,
    Index: "my_index",
    BatchSize: 10,
    DocumentToFields: func(ctx context.Context, doc *schema.Document) (map[string]FieldValue, error) {
        // 实现文档到字段的映射逻辑
        return fields, nil
    },
    ValidationMode: ValidationModeWarn,
    EnableSchemaCheck: true,
}

indexer, err := NewIndexer(ctx, config)
if err != nil {
    log.Fatal(err)
}
```

## 注意事项

1. 批量大小（BatchSize）建议根据实际需求设置，但不要超过 1000
2. 如果使用向量搜索功能，必须提供 Embedding 实现
3. 建议在生产环境中启用 EnableSchemaCheck 以确保数据一致性
4. 使用 ValidationModeWarn 可以在开发阶段更灵活地处理映射问题
5. 在开发阶段建议使用 ValidationModeWarn 模式，以便及时发现并处理验证问题
6. 生产环境建议使用 ValidationModeError 模式，以确保数据质量和一致性
