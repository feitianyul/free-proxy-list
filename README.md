**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-03-02 09:44:15 UTC（2026-03-02 17:44:15 UTC+8）

**代理总数：59**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 59 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 59 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 103.84.95.54:7890 | ✓ 838ms | 否 | ✓ 897ms | ✓ 1231ms | 否 | http |
| 205.209.118.30:3138 | ✓ 318ms | 否 | ✓ 673ms | ✓ 1286ms | 否 | http |
| 46.249.103.192:443 | ✓ 546ms | 否 | ✓ 1231ms | ✓ 1852ms | 否 | http |
| 5.75.196.26:40000 | ✓ 1531ms | ✓ 1420ms | ✓ 1610ms | ✓ 1835ms | ✓ 1084ms | http |
| 115.231.181.40:8128 | 否 | 否 | ✓ 1507ms | ✓ 1355ms | ✓ 1461ms | http |
| 59.46.216.131:30001 | ✓ 1290ms | ✓ 1577ms | ✓ 1263ms | ✓ 1669ms | 否 | http |
| 47.77.180.205:1080 | ✓ 915ms | ✓ 1598ms | ✓ 563ms | ✓ 967ms | 否 | http |
| 35.234.17.221:8080 | ✓ 1128ms | ✓ 1740ms | ✓ 1081ms | 否 | 否 | http |
| 120.132.97.88:7897 | ✓ 1599ms | 否 | ✓ 1435ms | 否 | ✓ 1029ms | http |
| 2.56.178.131:443 | ✓ 1091ms | 否 | ✓ 1454ms | ✓ 1716ms | 否 | http |
| 81.70.169.194:80 | ✓ 1212ms | ✓ 1545ms | ✓ 1094ms | ✓ 1451ms | ✓ 1267ms | http |
| 121.237.181.137:8888 | ✓ 1105ms | 否 | ✓ 1097ms | ✓ 1496ms | ✓ 985ms | http |
| 120.92.212.16:8890 | ✓ 1272ms | 否 | ✓ 1386ms | ✓ 1472ms | ✓ 1125ms | http |
| 101.43.255.96:80 | 否 | ✓ 1444ms | ✓ 1169ms | ✓ 1566ms | 否 | http |
| 35.225.22.61:80 | 否 | ✓ 1454ms | ✓ 372ms | ✓ 1306ms | 否 | http |
| 45.125.67.37:8443 | ✓ 1056ms | 否 | 否 | ✓ 1482ms | ✓ 1295ms | http |
| 165.227.5.10:8888 | 否 | ✓ 1354ms | ✓ 1164ms | 否 | ✓ 1575ms | http |
| 125.128.12.114:3128 | 否 | 否 | ✓ 1365ms | ✓ 1587ms | ✓ 993ms | http |
| 37.187.109.70:10111 | ✓ 1553ms | ✓ 1472ms | ✓ 465ms | ✓ 1571ms | ✓ 1639ms | http |
| 43.153.28.68:3128 | ✓ 518ms | 否 | ✓ 625ms | ✓ 987ms | 否 | http |
| 150.107.140.238:3128 | ✓ 1944ms | 否 | ✓ 1674ms | ✓ 1561ms | ✓ 1147ms | http |
| 171.234.62.116:10001 | ✓ 1921ms | 否 | ✓ 1936ms | 否 | ✓ 1481ms | http |
| 91.233.223.147:3128 | ✓ 1458ms | 否 | ✓ 711ms | ✓ 1822ms | ✓ 1468ms | http |
| 113.176.92.71:3128 | 否 | 否 | ✓ 1455ms | ✓ 1436ms | ✓ 1176ms | http |
| 103.139.138.194:3128 | 否 | 否 | ✓ 1600ms | ✓ 1618ms | ✓ 1281ms | http |
| 120.92.212.16:7890 | 否 | 否 | ✓ 1154ms | ✓ 1748ms | ✓ 1577ms | http |
| 34.101.184.164:3128 | ✓ 1303ms | 否 | ✓ 1660ms | ✓ 1594ms | ✓ 1327ms | http |
| 74.208.234.198:443 | 否 | ✓ 1783ms | ✓ 549ms | ✓ 1104ms | ✓ 1352ms | http |
| 45.81.131.145:8888 | ✓ 997ms | 否 | ✓ 1256ms | ✓ 1256ms | ✓ 827ms | http |
| 103.215.36.88:18127 | ✓ 1143ms | ✓ 1469ms | ✓ 1558ms | ✓ 1768ms | ✓ 1137ms | http |
| 20.2.83.243:3128 | ✓ 1038ms | 否 | ✓ 1044ms | ✓ 1022ms | ✓ 897ms | http |
| 185.243.218.43:49153 | ✓ 591ms | ✓ 1576ms | ✓ 1462ms | ✓ 1875ms | ✓ 1583ms | http |
| 121.230.9.113:1080 | ✓ 1322ms | 否 | ✓ 1335ms | 否 | ✓ 1196ms | http |
| 47.84.131.156:8100 | 否 | 否 | ✓ 901ms | ✓ 1320ms | ✓ 1053ms | http |
| 85.198.84.77:10808 | ✓ 1138ms | 否 | ✓ 1670ms | 否 | ✓ 1817ms | http |
| 45.136.198.40:3128 | ✓ 1681ms | 否 | ✓ 1532ms | ✓ 1920ms | ✓ 1543ms | http |
| 111.79.111.126:3128 | ✓ 1245ms | 否 | 否 | ✓ 1869ms | ✓ 1956ms | http |
| 45.140.147.82:1081 | ✓ 375ms | ✓ 1041ms | ✓ 1620ms | ✓ 1715ms | ✓ 1379ms | http |
| 207.254.71.62:8088 | ✓ 851ms | ✓ 1961ms | ✓ 1515ms | 否 | ✓ 1688ms | http |
| 115.76.5.32:10008 | ✓ 1800ms | 否 | ✓ 1923ms | 否 | ✓ 1725ms | http |
| 36.147.78.166:443 | ✓ 1916ms | 否 | 否 | ✓ 1821ms | ✓ 1833ms | http |
| 103.215.36.88:15852 | ✓ 1205ms | 否 | ✓ 1205ms | 否 | ✓ 1243ms | http |
| 103.215.36.88:10101 | ✓ 1129ms | 否 | ✓ 1500ms | ✓ 1525ms | ✓ 1110ms | http |
| 103.82.23.118:5247 | ✓ 1624ms | 否 | ✓ 1389ms | ✓ 1892ms | ✓ 1416ms | http |
| 57.128.188.167:9174 | ✓ 1938ms | 否 | ✓ 1614ms | ✓ 1879ms | 否 | http |
| 121.230.8.136:1080 | ✓ 1730ms | ✓ 1976ms | ✓ 1661ms | 否 | 否 | http |
| 171.234.62.116:10003 | 否 | 否 | ✓ 1604ms | ✓ 1765ms | ✓ 1498ms | http |
| 221.122.91.36:11273 | ✓ 1117ms | ✓ 1419ms | ✓ 1071ms | ✓ 1395ms | ✓ 1101ms | http |
| 221.122.91.36:11195 | ✓ 1189ms | ✓ 1406ms | ✓ 1080ms | ✓ 1435ms | ✓ 1078ms | http |
| 103.39.51.190:8080 | ✓ 1567ms | 否 | 否 | ✓ 1652ms | ✓ 1645ms | http |
| 5.129.228.225:1080 | ✓ 1232ms | 否 | 否 | ✓ 1238ms | ✓ 1911ms | http |
| 192.71.213.85:9812 | ✓ 392ms | 否 | ✓ 392ms | ✓ 1414ms | 否 | http |
| 183.128.208.68:7890 | ✓ 1911ms | 否 | ✓ 1723ms | ✓ 1446ms | ✓ 1194ms | http |
| 121.230.9.168:1080 | 否 | ✓ 1568ms | ✓ 1588ms | ✓ 1851ms | ✓ 1843ms | http |
| 172.212.68.37:3128 | ✓ 1213ms | 否 | ✓ 811ms | ✓ 1380ms | ✓ 1347ms | http |
| 45.186.6.104:3128 | ✓ 996ms | ✓ 1841ms | ✓ 1644ms | 否 | 否 | http |
| 113.255.59.226:8080 | ✓ 1789ms | 否 | ✓ 1703ms | ✓ 1352ms | ✓ 1411ms | http |
| 20.120.225.109:3128 | ✓ 1263ms | 否 | ✓ 1150ms | ✓ 1616ms | ✓ 1005ms | http |
| 120.55.163.237:10086 | ✓ 1018ms | ✓ 1297ms | ✓ 1454ms | 否 | 否 | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
