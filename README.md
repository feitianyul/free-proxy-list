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

最后更新：2026-02-24 11:48:10 UTC（2026-02-24 19:48:10 UTC+8）

**代理总数：73**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 73 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 73 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 1013ms | 否 | ✓ 1803ms | ✓ 1125ms | ✓ 1084ms | http |
| 35.225.22.61:80 | ✓ 1208ms | ✓ 1090ms | ✓ 418ms | ✓ 1258ms | ✓ 795ms | http |
| 45.151.182.9:3128 | ✓ 888ms | 否 | ✓ 1185ms | 否 | ✓ 1832ms | http |
| 211.230.49.122:3128 | ✓ 1179ms | ✓ 1421ms | 否 | ✓ 1328ms | 否 | http |
| 72.56.59.56:63127 | ✓ 1604ms | 否 | ✓ 1585ms | 否 | ✓ 1838ms | http |
| 72.56.59.17:61931 | ✓ 1606ms | 否 | ✓ 1615ms | 否 | ✓ 1935ms | http |
| 202.152.44.19:8081 | ✓ 1560ms | 否 | ✓ 1269ms | ✓ 1349ms | 否 | http |
| 132.145.93.138:1080 | ✓ 1701ms | 否 | 否 | ✓ 1698ms | ✓ 1938ms | http |
| 202.152.44.18:8081 | ✓ 1113ms | 否 | ✓ 1081ms | ✓ 1459ms | ✓ 1226ms | http |
| 104.238.30.63:63744 | ✓ 1712ms | 否 | ✓ 1870ms | 否 | ✓ 1999ms | http |
| 125.128.12.94:3128 | ✓ 1494ms | ✓ 1236ms | 否 | 否 | ✓ 1475ms | http |
| 171.229.213.253:10019 | ✓ 1515ms | 否 | ✓ 1612ms | ✓ 1713ms | ✓ 1567ms | http |
| 150.107.140.238:3128 | ✓ 1631ms | 否 | ✓ 1837ms | ✓ 1797ms | ✓ 1765ms | http |
| 190.242.157.215:8080 | ✓ 1241ms | 否 | ✓ 790ms | ✓ 1752ms | 否 | http |
| 144.124.253.249:47561 | ✓ 1066ms | ✓ 1461ms | ✓ 1571ms | ✓ 1899ms | 否 | http |
| 58.69.185.88:8080 | ✓ 1890ms | 否 | ✓ 1557ms | ✓ 1660ms | ✓ 1861ms | http |
| 120.92.212.16:7890 | ✓ 1133ms | 否 | ✓ 1168ms | ✓ 1683ms | ✓ 1090ms | http |
| 34.50.41.78:8888 | ✓ 1725ms | 否 | ✓ 1322ms | ✓ 1323ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1213ms | ✓ 1506ms | ✓ 1351ms | ✓ 1384ms | ✓ 1087ms | http |
| 72.56.59.62:63133 | ✓ 1513ms | 否 | ✓ 1654ms | 否 | ✓ 1935ms | http |
| 81.70.169.194:80 | ✓ 1072ms | ✓ 1512ms | ✓ 1663ms | ✓ 1480ms | 否 | http |
| 101.43.255.96:80 | ✓ 1181ms | ✓ 1484ms | ✓ 1172ms | 否 | ✓ 1155ms | http |
| 18.229.201.117:3128 | ✓ 590ms | ✓ 1928ms | 否 | 否 | ✓ 1538ms | http |
| 52.188.28.218:3128 | 否 | 否 | ✓ 1469ms | ✓ 1086ms | ✓ 1747ms | http |
| 35.234.17.221:8080 | 否 | 否 | ✓ 1288ms | ✓ 1174ms | ✓ 1155ms | http |
| 103.84.95.54:7890 | ✓ 1162ms | 否 | ✓ 841ms | 否 | ✓ 780ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1335ms | ✓ 1097ms | ✓ 1270ms | ✓ 1884ms | http |
| 14.56.107.244:3128 | ✓ 1732ms | ✓ 1496ms | 否 | ✓ 1179ms | 否 | http |
| 72.56.50.17:59787 | ✓ 1662ms | 否 | ✓ 1935ms | 否 | ✓ 1836ms | http |
| 62.113.119.14:8080 | ✓ 616ms | ✓ 1639ms | ✓ 949ms | ✓ 1511ms | ✓ 1199ms | http |
| 144.31.69.170:1080 | ✓ 1040ms | 否 | ✓ 1278ms | 否 | ✓ 1650ms | http |
| 190.9.109.199:999 | ✓ 742ms | 否 | ✓ 1122ms | ✓ 1476ms | ✓ 1091ms | http |
| 190.9.109.205:999 | ✓ 674ms | 否 | ✓ 1187ms | ✓ 1495ms | ✓ 1134ms | http |
| 45.140.147.82:1082 | ✓ 416ms | ✓ 1918ms | ✓ 734ms | 否 | ✓ 982ms | http |
| 45.140.147.82:1081 | ✓ 412ms | ✓ 1628ms | ✓ 1030ms | ✓ 1865ms | ✓ 1119ms | http |
| 104.238.30.50:59741 | ✓ 1803ms | 否 | ✓ 1832ms | 否 | ✓ 1999ms | http |
| 59.46.216.131:30001 | ✓ 1235ms | ✓ 1607ms | 否 | ✓ 1592ms | 否 | http |
| 147.45.159.213:48206 | ✓ 887ms | 否 | ✓ 1867ms | 否 | ✓ 1831ms | http |
| 36.136.27.2:4999 | ✓ 1481ms | ✓ 1455ms | ✓ 1329ms | 否 | ✓ 1238ms | http |
| 47.101.159.19:8899 | ✓ 993ms | ✓ 1155ms | ✓ 1005ms | 否 | 否 | http |
| 91.238.104.171:2023 | ✓ 1086ms | 否 | ✓ 1979ms | ✓ 1946ms | 否 | http |
| 124.16.93.70:7890 | ✓ 1039ms | ✓ 1573ms | ✓ 1145ms | ✓ 1375ms | ✓ 1069ms | http |
| 106.14.205.114:483 | 否 | ✓ 1518ms | ✓ 1866ms | ✓ 1190ms | ✓ 964ms | http |
| 103.82.93.100:3128 | 否 | 否 | ✓ 1494ms | ✓ 1373ms | ✓ 1068ms | http |
| 187.216.141.46:3128 | ✓ 1034ms | ✓ 1994ms | 否 | ✓ 1859ms | 否 | http |
| 37.27.100.102:443 | ✓ 1902ms | ✓ 1636ms | 否 | ✓ 1665ms | 否 | http |
| 120.46.152.136:3128 | ✓ 1280ms | ✓ 1579ms | ✓ 1658ms | 否 | 否 | http |
| 138.197.68.35:4857 | 否 | 否 | ✓ 901ms | ✓ 1010ms | ✓ 942ms | http |
| 162.240.154.26:3128 | ✓ 885ms | 否 | ✓ 866ms | ✓ 935ms | ✓ 743ms | http |
| 137.220.150.22:6005 | ✓ 1070ms | 否 | ✓ 1309ms | ✓ 1651ms | ✓ 965ms | http |
| 104.238.30.38:59741 | ✓ 1844ms | 否 | ✓ 1801ms | 否 | ✓ 1999ms | http |
| 78.13.231.158:3128 | ✓ 472ms | 否 | ✓ 1655ms | ✓ 1888ms | ✓ 1496ms | http |
| 18.229.170.122:3128 | ✓ 707ms | 否 | ✓ 798ms | 否 | ✓ 1601ms | http |
| 211.171.114.154:3128 | ✓ 1288ms | 否 | ✓ 1434ms | ✓ 1694ms | ✓ 1664ms | http |
| 72.56.59.23:61937 | ✓ 1867ms | 否 | ✓ 1644ms | 否 | ✓ 1909ms | http |
| 36.147.78.166:80 | ✓ 1952ms | ✓ 1773ms | ✓ 1870ms | 否 | 否 | http |
| 176.124.220.172:3128 | ✓ 615ms | 否 | ✓ 1556ms | ✓ 1935ms | ✓ 1677ms | http |
| 104.238.30.37:59741 | ✓ 1775ms | 否 | ✓ 1942ms | 否 | ✓ 1992ms | http |
| 178.238.117.178:8080 | ✓ 929ms | 否 | ✓ 1023ms | 否 | ✓ 1847ms | http |
| 202.129.206.239:3128 | ✓ 1530ms | 否 | ✓ 1423ms | ✓ 1695ms | 否 | http |
| 104.238.30.39:59741 | ✓ 1718ms | 否 | ✓ 1843ms | 否 | ✓ 1995ms | http |
| 104.238.30.40:59741 | ✓ 1768ms | 否 | ✓ 1807ms | 否 | ✓ 1998ms | http |
| 217.216.109.116:8080 | ✓ 889ms | 否 | ✓ 1976ms | ✓ 1476ms | ✓ 1852ms | http |
| 113.45.250.180:443 | 否 | ✓ 1390ms | 否 | ✓ 1304ms | ✓ 1074ms | http |
| 193.181.35.35:8118 | ✓ 804ms | 否 | ✓ 1172ms | ✓ 1929ms | ✓ 1688ms | http |
| 20.120.225.109:3128 | ✓ 1068ms | 否 | ✓ 808ms | ✓ 1185ms | ✓ 1032ms | http |
| 138.124.53.25:7443 | ✓ 435ms | 否 | ✓ 1557ms | ✓ 1310ms | ✓ 1105ms | http |
| 45.12.151.226:2828 | 否 | 否 | ✓ 1663ms | ✓ 1997ms | ✓ 1660ms | http |
| 103.39.51.190:8080 | ✓ 1444ms | 否 | 否 | ✓ 1511ms | ✓ 1693ms | http |
| 165.227.5.10:8888 | 否 | 否 | ✓ 497ms | ✓ 1111ms | ✓ 1059ms | http |
| 103.35.188.243:3128 | 否 | ✓ 1670ms | 否 | ✓ 1897ms | ✓ 1851ms | http |
| 91.233.223.147:3128 | ✓ 1262ms | 否 | ✓ 1460ms | ✓ 1978ms | 否 | http |
| 178.130.47.129:1082 | ✓ 1156ms | 否 | ✓ 795ms | ✓ 890ms | 否 | http |

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
