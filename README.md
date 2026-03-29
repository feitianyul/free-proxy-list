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

最后更新：2026-03-29 15:25:53 UTC（2026-03-29 23:25:53 UTC+8）

**代理总数：82**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 81 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 82 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 39.185.46.193:5911 | ✓ 780ms | ✓ 963ms | ✓ 927ms | ✓ 1050ms | ✓ 811ms | http |
| 43.99.54.236:5555 | ✓ 1043ms | 否 | ✓ 761ms | ✓ 992ms | ✓ 902ms | http |
| 147.161.210.140:8800 | ✓ 1662ms | 否 | ✓ 918ms | ✓ 1168ms | ✓ 1177ms | http |
| 219.117.204.211:7799 | ✓ 1661ms | 否 | ✓ 1061ms | ✓ 1117ms | ✓ 1126ms | http |
| 167.103.115.102:8800 | ✓ 1864ms | 否 | ✓ 1170ms | 否 | ✓ 1297ms | http |
| 35.225.22.61:80 | 否 | ✓ 1290ms | ✓ 1876ms | ✓ 956ms | 否 | http |
| 1.231.81.166:3128 | ✓ 928ms | ✓ 1331ms | ✓ 1616ms | ✓ 1236ms | ✓ 944ms | http |
| 167.103.34.108:8800 | ✓ 1241ms | 否 | ✓ 1261ms | ✓ 1453ms | 否 | http |
| 103.84.95.54:7890 | ✓ 1073ms | 否 | ✓ 848ms | 否 | ✓ 847ms | http |
| 167.103.144.127:8800 | ✓ 1676ms | 否 | ✓ 1347ms | 否 | ✓ 1603ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1671ms | ✓ 1785ms | ✓ 1424ms | 否 | http |
| 167.103.31.122:8800 | ✓ 1948ms | 否 | ✓ 1667ms | 否 | ✓ 1907ms | http |
| 45.167.124.52:8080 | ✓ 1722ms | 否 | ✓ 740ms | ✓ 1516ms | ✓ 1324ms | http |
| 137.184.1.87:3128 | ✓ 1211ms | 否 | ✓ 1490ms | ✓ 888ms | ✓ 660ms | http |
| 42.96.16.158:1311 | ✓ 1101ms | 否 | ✓ 1273ms | ✓ 1390ms | ✓ 1067ms | http |
| 95.213.217.168:52004 | 否 | 否 | ✓ 1005ms | ✓ 1607ms | ✓ 1197ms | http |
| 168.110.52.228:3128 | ✓ 1750ms | 否 | 否 | ✓ 1427ms | ✓ 864ms | http |
| 101.47.73.135:3128 | ✓ 1135ms | 否 | ✓ 1618ms | 否 | ✓ 1617ms | http |
| 20.210.39.153:8561 | ✓ 1597ms | 否 | 否 | ✓ 1018ms | ✓ 820ms | http |
| 147.161.239.240:8800 | ✓ 739ms | 否 | ✓ 1489ms | ✓ 1493ms | 否 | http |
| 101.43.127.100:8877 | ✓ 1032ms | ✓ 1223ms | ✓ 1020ms | 否 | ✓ 1884ms | http |
| 193.233.22.29:10808 | ✓ 1328ms | 否 | ✓ 879ms | ✓ 1641ms | 否 | http |
| 5.102.109.41:999 | ✓ 1320ms | ✓ 1958ms | ✓ 1047ms | ✓ 1646ms | ✓ 1210ms | http |
| 177.234.217.88:999 | ✓ 1408ms | 否 | ✓ 1735ms | 否 | ✓ 1581ms | http |
| 103.113.70.189:1081 | ✓ 1061ms | ✓ 1692ms | ✓ 102ms | ✓ 1833ms | ✓ 956ms | http |
| 137.184.6.37:3128 | ✓ 1082ms | 否 | ✓ 780ms | ✓ 843ms | ✓ 688ms | http |
| 62.234.206.73:3128 | ✓ 1436ms | ✓ 1402ms | 否 | ✓ 1996ms | ✓ 1845ms | http |
| 20.27.11.248:8561 | 否 | 否 | ✓ 578ms | ✓ 1213ms | ✓ 1621ms | http |
| 20.27.15.111:8561 | 否 | 否 | ✓ 610ms | ✓ 1268ms | ✓ 1651ms | http |
| 103.93.93.77:8050 | 否 | 否 | ✓ 1763ms | ✓ 1662ms | ✓ 1637ms | http |
| 120.92.212.16:8890 | ✓ 1114ms | ✓ 1341ms | ✓ 1387ms | ✓ 1375ms | ✓ 1124ms | http |
| 45.12.151.226:2829 | ✓ 778ms | 否 | ✓ 1532ms | 否 | ✓ 1455ms | http |
| 120.92.212.16:7890 | ✓ 1116ms | 否 | 否 | ✓ 1611ms | ✓ 1350ms | http |
| 45.140.147.155:1081 | ✓ 1636ms | 否 | ✓ 1419ms | ✓ 1508ms | ✓ 979ms | http |
| 106.75.15.167:7890 | 否 | 否 | ✓ 1224ms | ✓ 1321ms | ✓ 1076ms | http |
| 208.87.243.199:7878 | ✓ 1632ms | 否 | 否 | ✓ 1146ms | ✓ 717ms | http |
| 222.228.171.92:8080 | ✓ 1922ms | 否 | 否 | ✓ 1615ms | ✓ 1255ms | http |
| 38.34.179.67:8451 | ✓ 463ms | ✓ 1249ms | ✓ 1440ms | ✓ 952ms | ✓ 1420ms | http |
| 24.144.86.173:1080 | ✓ 1053ms | ✓ 1646ms | ✓ 1556ms | ✓ 1449ms | 否 | http |
| 20.78.26.206:8561 | 否 | 否 | ✓ 607ms | ✓ 926ms | ✓ 737ms | http |
| 20.78.118.91:8561 | 否 | 否 | ✓ 629ms | ✓ 897ms | ✓ 743ms | http |
| 38.145.218.229:8453 | ✓ 1185ms | ✓ 867ms | ✓ 429ms | ✓ 1253ms | ✓ 1920ms | http |
| 20.27.13.35:8561 | ✓ 1458ms | ✓ 1443ms | ✓ 770ms | ✓ 1029ms | ✓ 900ms | http |
| 20.27.14.220:8561 | ✓ 1458ms | ✓ 1361ms | ✓ 734ms | ✓ 1078ms | ✓ 968ms | http |
| 38.34.179.106:8446 | ✓ 1081ms | ✓ 1825ms | ✓ 1149ms | 否 | 否 | http |
| 217.76.245.80:999 | 否 | ✓ 1469ms | ✓ 1289ms | 否 | ✓ 1231ms | http |
| 101.32.244.83:8080 | ✓ 1078ms | 否 | ✓ 1074ms | ✓ 1618ms | ✓ 1452ms | http |
| 137.184.1.155:3128 | ✓ 1203ms | 否 | ✓ 1586ms | ✓ 1317ms | ✓ 880ms | http |
| 88.80.150.82:8080 | ✓ 880ms | ✓ 1658ms | 否 | ✓ 1944ms | ✓ 1647ms | https |
| 38.145.208.235:8445 | 否 | ✓ 1485ms | ✓ 833ms | ✓ 1311ms | ✓ 882ms | http |
| 20.210.76.104:8561 | ✓ 1383ms | ✓ 1295ms | ✓ 1003ms | ✓ 1182ms | ✓ 1009ms | http |
| 38.145.208.246:8450 | 否 | ✓ 1093ms | ✓ 1225ms | ✓ 1619ms | ✓ 885ms | http |
| 38.34.179.35:8448 | 否 | 否 | ✓ 971ms | ✓ 1272ms | ✓ 902ms | http |
| 51.79.207.21:8080 | ✓ 1094ms | 否 | 否 | ✓ 1595ms | ✓ 1641ms | http |
| 45.136.131.38:8445 | ✓ 273ms | ✓ 1272ms | ✓ 826ms | ✓ 1210ms | ✓ 825ms | http |
| 38.145.208.245:8452 | ✓ 266ms | ✓ 1937ms | ✓ 263ms | ✓ 1312ms | ✓ 671ms | http |
| 38.145.218.228:8452 | ✓ 347ms | ✓ 1958ms | ✓ 843ms | ✓ 1155ms | ✓ 713ms | http |
| 38.145.208.151:8453 | ✓ 1506ms | 否 | ✓ 497ms | ✓ 909ms | ✓ 675ms | http |
| 38.34.179.17:8453 | ✓ 650ms | ✓ 1131ms | ✓ 1926ms | ✓ 988ms | ✓ 903ms | http |
| 38.34.179.12:8452 | ✓ 740ms | ✓ 1793ms | ✓ 1171ms | ✓ 1037ms | ✓ 869ms | http |
| 45.136.131.42:8452 | ✓ 356ms | ✓ 1109ms | ✓ 899ms | ✓ 1218ms | ✓ 1006ms | http |
| 38.145.208.171:8453 | ✓ 1593ms | ✓ 1016ms | ✓ 260ms | 否 | ✓ 1472ms | http |
| 45.136.130.184:8453 | 否 | 否 | ✓ 308ms | ✓ 1102ms | ✓ 1063ms | http |
| 38.145.208.193:8451 | 否 | ✓ 1861ms | ✓ 569ms | ✓ 1264ms | ✓ 1188ms | http |
| 38.145.220.198:8451 | ✓ 1985ms | 否 | ✓ 734ms | ✓ 889ms | ✓ 901ms | http |
| 38.145.203.132:8450 | 否 | 否 | ✓ 855ms | ✓ 1547ms | ✓ 645ms | http |
| 38.34.179.202:8449 | 否 | 否 | ✓ 352ms | ✓ 1038ms | ✓ 897ms | http |
| 38.34.183.224:8451 | 否 | 否 | ✓ 1610ms | ✓ 1689ms | ✓ 760ms | http |
| 38.34.179.95:8444 | ✓ 1470ms | 否 | ✓ 1406ms | ✓ 1038ms | ✓ 1929ms | http |
| 85.208.108.43:10808 | ✓ 986ms | 否 | ✓ 499ms | ✓ 1374ms | ✓ 680ms | http |
| 85.208.108.43:2094 | 否 | 否 | ✓ 515ms | ✓ 982ms | ✓ 681ms | http |
| 38.34.179.164:8448 | ✓ 1627ms | ✓ 1532ms | ✓ 811ms | 否 | ✓ 733ms | http |
| 45.136.198.40:3128 | ✓ 1568ms | ✓ 1576ms | 否 | ✓ 1916ms | ✓ 1610ms | http |
| 38.145.203.135:8444 | 否 | ✓ 953ms | ✓ 633ms | ✓ 1534ms | ✓ 1727ms | http |
| 64.227.76.27:1080 | ✓ 540ms | 否 | ✓ 1605ms | ✓ 1705ms | 否 | http |
| 45.144.28.81:10808 | ✓ 456ms | 否 | ✓ 1185ms | ✓ 1434ms | 否 | http |
| 121.43.196.213:8222 | ✓ 1030ms | ✓ 1209ms | ✓ 982ms | ✓ 1240ms | ✓ 994ms | http |
| 121.43.196.210:8222 | ✓ 1063ms | 否 | ✓ 945ms | ✓ 1322ms | ✓ 963ms | http |
| 114.55.226.123:10086 | ✓ 1140ms | 否 | ✓ 1097ms | ✓ 1404ms | ✓ 1134ms | http |
| 103.39.51.190:8080 | ✓ 1534ms | 否 | 否 | ✓ 1813ms | ✓ 1589ms | http |
| 61.52.131.172:8443 | ✓ 950ms | ✓ 1243ms | ✓ 1105ms | ✓ 1298ms | ✓ 1017ms | http |
| 116.80.65.80:3172 | ✓ 1761ms | 否 | ✓ 1707ms | ✓ 1990ms | ✓ 1764ms | http |

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
