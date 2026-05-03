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

最后更新：2026-05-03 04:29:56 UTC（2026-05-03 12:29:56 UTC+8）

**代理总数：66**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 66 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 66 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 149.51.42.10:3128 | ✓ 522ms | ✓ 1271ms | 否 | ✓ 1445ms | 否 | http |
| 47.85.51.197:1080 | ✓ 95ms | ✓ 1912ms | 否 | 否 | ✓ 955ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1610ms | ✓ 1300ms | 否 | ✓ 1132ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1631ms | ✓ 1685ms | ✓ 1531ms | ✓ 1238ms | http |
| 38.180.62.47:10808 | ✓ 1816ms | ✓ 1200ms | ✓ 1354ms | 否 | 否 | http |
| 148.230.4.241:999 | ✓ 924ms | ✓ 1910ms | ✓ 834ms | ✓ 1803ms | ✓ 1454ms | http |
| 45.167.124.71:999 | ✓ 1320ms | 否 | ✓ 1393ms | 否 | ✓ 1658ms | http |
| 45.140.147.155:1082 | ✓ 443ms | ✓ 1253ms | ✓ 1055ms | ✓ 1720ms | ✓ 1261ms | http |
| 218.108.131.186:17890 | ✓ 1011ms | ✓ 1281ms | ✓ 953ms | ✓ 1269ms | ✓ 1087ms | http |
| 62.133.60.126:24558 | 否 | 否 | ✓ 1167ms | ✓ 1767ms | ✓ 1825ms | http |
| 103.3.246.71:3128 | ✓ 1617ms | 否 | ✓ 1311ms | ✓ 1602ms | ✓ 1142ms | http |
| 45.153.231.229:8080 | ✓ 1317ms | ✓ 1700ms | ✓ 980ms | ✓ 1954ms | ✓ 1799ms | http |
| 147.45.178.211:14658 | ✓ 1331ms | 否 | ✓ 1137ms | ✓ 1535ms | 否 | http |
| 45.59.122.132:80 | ✓ 752ms | ✓ 1844ms | ✓ 758ms | ✓ 1220ms | ✓ 1452ms | http |
| 212.58.132.5:8888 | ✓ 1053ms | 否 | ✓ 1092ms | ✓ 1523ms | ✓ 1174ms | http |
| 20.78.118.91:8561 | ✓ 1604ms | ✓ 1156ms | ✓ 643ms | ✓ 994ms | ✓ 822ms | http |
| 20.210.39.153:8561 | ✓ 1592ms | ✓ 1327ms | ✓ 639ms | ✓ 1006ms | ✓ 795ms | http |
| 20.78.26.206:8561 | ✓ 1594ms | 否 | ✓ 644ms | ✓ 992ms | ✓ 760ms | http |
| 94.131.118.129:1081 | ✓ 1490ms | 否 | ✓ 910ms | 否 | ✓ 1143ms | http |
| 45.125.67.37:8443 | ✓ 1265ms | 否 | ✓ 1889ms | ✓ 1647ms | ✓ 1263ms | http |
| 3.101.133.120:80 | ✓ 820ms | ✓ 1473ms | ✓ 1576ms | ✓ 1481ms | ✓ 1300ms | http |
| 109.120.156.122:8090 | ✓ 848ms | ✓ 1987ms | ✓ 1546ms | 否 | ✓ 1570ms | http |
| 130.61.174.200:1080 | ✓ 651ms | ✓ 1536ms | 否 | ✓ 1328ms | 否 | http |
| 20.27.13.35:8561 | 否 | 否 | ✓ 1764ms | ✓ 1870ms | ✓ 1471ms | http |
| 59.46.216.131:30001 | ✓ 1227ms | ✓ 1565ms | ✓ 1674ms | 否 | 否 | http |
| 38.180.121.135:10808 | 否 | ✓ 1584ms | ✓ 1124ms | ✓ 1571ms | 否 | http |
| 46.105.190.38:3128 | ✓ 585ms | ✓ 1439ms | ✓ 437ms | 否 | 否 | http |
| 206.206.126.177:2412 | ✓ 1504ms | ✓ 1725ms | ✓ 974ms | ✓ 1225ms | ✓ 965ms | http |
| 49.156.44.114:8080 | ✓ 1628ms | ✓ 1803ms | ✓ 1616ms | ✓ 1702ms | ✓ 1657ms | http |
| 20.27.11.248:8561 | 否 | ✓ 1773ms | ✓ 696ms | ✓ 1028ms | ✓ 850ms | http |
| 20.27.14.220:8561 | 否 | ✓ 1849ms | ✓ 662ms | ✓ 1008ms | ✓ 822ms | http |
| 20.27.15.111:8561 | 否 | ✓ 1683ms | ✓ 683ms | ✓ 1078ms | ✓ 906ms | http |
| 103.157.200.126:3128 | ✓ 1568ms | 否 | ✓ 1329ms | ✓ 1638ms | ✓ 1319ms | http |
| 103.35.190.69:1082 | ✓ 208ms | ✓ 1212ms | ✓ 169ms | ✓ 1217ms | ✓ 731ms | http |
| 149.51.42.10:8080 | ✓ 393ms | ✓ 1264ms | 否 | ✓ 1235ms | 否 | http |
| 139.5.189.229:8888 | ✓ 1238ms | 否 | ✓ 1427ms | ✓ 1595ms | ✓ 1165ms | http |
| 190.12.150.244:999 | ✓ 1515ms | 否 | ✓ 943ms | ✓ 1637ms | ✓ 1350ms | http |
| 86.104.74.110:1081 | ✓ 1506ms | ✓ 1616ms | ✓ 350ms | ✓ 1579ms | 否 | http |
| 86.104.72.220:1081 | ✓ 297ms | ✓ 967ms | ✓ 390ms | ✓ 1631ms | 否 | http |
| 86.104.74.110:1082 | ✓ 531ms | ✓ 1199ms | ✓ 697ms | ✓ 1725ms | ✓ 1140ms | http |
| 34.96.238.40:8080 | 否 | 否 | ✓ 1093ms | ✓ 1835ms | ✓ 1628ms | http |
| 185.230.191.240:3128 | ✓ 651ms | ✓ 1610ms | ✓ 838ms | 否 | ✓ 1596ms | http |
| 150.107.140.238:3128 | ✓ 1281ms | 否 | ✓ 1944ms | 否 | ✓ 1843ms | http |
| 193.123.250.39:1080 | ✓ 1880ms | 否 | ✓ 1225ms | ✓ 1260ms | ✓ 1032ms | http |
| 8.154.21.175:3128 | ✓ 1085ms | ✓ 1224ms | ✓ 1027ms | ✓ 1362ms | ✓ 1082ms | http |
| 154.90.48.209:9090 | ✓ 1004ms | 否 | ✓ 1328ms | ✓ 1634ms | ✓ 1676ms | http |
| 101.32.243.189:80 | ✓ 1743ms | 否 | ✓ 1526ms | ✓ 1470ms | ✓ 1483ms | http |
| 152.42.177.32:8888 | ✓ 1746ms | 否 | ✓ 1946ms | ✓ 1442ms | ✓ 1521ms | http |
| 91.233.223.147:3128 | ✓ 1027ms | 否 | ✓ 1608ms | 否 | ✓ 1483ms | http |
| 94.131.118.129:1082 | ✓ 977ms | ✓ 1450ms | ✓ 601ms | ✓ 1616ms | ✓ 1105ms | http |
| 119.195.17.15:3180 | 否 | ✓ 1814ms | ✓ 1814ms | ✓ 1516ms | ✓ 1250ms | http |
| 217.182.195.221:30001 | ✓ 992ms | 否 | 否 | ✓ 1735ms | ✓ 1850ms | http |
| 121.135.144.141:8089 | ✓ 1802ms | 否 | ✓ 1406ms | ✓ 1215ms | ✓ 1212ms | http |
| 152.70.91.193:40000 | ✓ 1710ms | 否 | 否 | ✓ 1567ms | ✓ 1179ms | http |
| 20.164.75.153:8080 | ✓ 1777ms | 否 | 否 | ✓ 1936ms | ✓ 1758ms | http |
| 144.124.227.88:3128 | ✓ 572ms | 否 | ✓ 1208ms | ✓ 1964ms | 否 | http |
| 45.88.0.116:3128 | ✓ 947ms | ✓ 1579ms | 否 | ✓ 1863ms | 否 | http |
| 45.129.141.143:3128 | ✓ 1856ms | 否 | ✓ 1791ms | 否 | ✓ 1880ms | http |
| 121.230.8.136:1080 | ✓ 1222ms | ✓ 1581ms | 否 | ✓ 1897ms | ✓ 1186ms | http |
| 210.223.44.230:3128 | ✓ 1412ms | 否 | 否 | ✓ 1770ms | ✓ 1372ms | http |
| 77.110.107.80:8080 | ✓ 443ms | 否 | ✓ 1662ms | ✓ 1954ms | 否 | http |
| 103.35.190.69:1081 | ✓ 232ms | ✓ 1512ms | ✓ 62ms | ✓ 928ms | ✓ 912ms | http |
| 5.253.43.103:3128 | ✓ 705ms | ✓ 1393ms | ✓ 917ms | ✓ 1650ms | ✓ 1272ms | http |
| 61.52.131.172:8443 | ✓ 1067ms | ✓ 1383ms | ✓ 1177ms | ✓ 1450ms | ✓ 1173ms | http |
| 106.10.55.212:1121 | ✓ 1771ms | ✓ 1102ms | 否 | ✓ 1418ms | ✓ 1148ms | http |
| 80.92.204.47:1081 | ✓ 1093ms | ✓ 1390ms | ✓ 669ms | ✓ 1477ms | ✓ 1118ms | http |

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
