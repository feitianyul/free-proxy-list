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

最后更新：2026-03-01 19:36:00 UTC（2026-03-02 03:36:00 UTC+8）

**代理总数：87**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 87 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 87 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | ✓ 456ms | 否 | 否 | ✓ 997ms | ✓ 653ms | http |
| 141.11.210.35:1080 | ✓ 300ms | ✓ 1625ms | ✓ 1616ms | ✓ 1143ms | ✓ 968ms | http |
| 183.249.5.109:22222 | 否 | ✓ 1063ms | ✓ 848ms | ✓ 1308ms | ✓ 960ms | http |
| 205.209.118.30:3138 | ✓ 303ms | ✓ 1626ms | ✓ 846ms | 否 | 否 | http |
| 74.208.234.198:443 | ✓ 708ms | 否 | 否 | ✓ 1426ms | ✓ 1046ms | http |
| 167.160.184.231:6005 | ✓ 1235ms | 否 | ✓ 1968ms | ✓ 1955ms | ✓ 1755ms | http |
| 14.56.107.244:3128 | ✓ 1269ms | 否 | 否 | ✓ 1173ms | ✓ 954ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1469ms | ✓ 1328ms | 否 | ✓ 1138ms | http |
| 186.148.180.46:999 | ✓ 1386ms | 否 | ✓ 1255ms | ✓ 1622ms | ✓ 1478ms | http |
| 148.135.85.87:1080 | ✓ 788ms | 否 | ✓ 1322ms | 否 | ✓ 842ms | http |
| 121.128.121.54:3128 | ✓ 1515ms | ✓ 1650ms | ✓ 744ms | 否 | ✓ 904ms | http |
| 104.238.30.50:59741 | ✓ 1774ms | 否 | ✓ 1807ms | 否 | ✓ 1971ms | http |
| 61.72.110.54:3128 | 否 | ✓ 1975ms | ✓ 1446ms | ✓ 1217ms | ✓ 1475ms | http |
| 190.9.109.196:999 | ✓ 1024ms | ✓ 1300ms | ✓ 1067ms | ✓ 1472ms | ✓ 1192ms | http |
| 104.238.30.45:59741 | ✓ 1774ms | 否 | ✓ 1904ms | 否 | ✓ 1997ms | http |
| 190.9.109.202:999 | ✓ 1025ms | ✓ 1607ms | ✓ 898ms | ✓ 1331ms | 否 | http |
| 120.238.159.230:22222 | 否 | 否 | ✓ 983ms | ✓ 1230ms | ✓ 1014ms | http |
| 120.92.212.16:7890 | ✓ 1298ms | ✓ 1314ms | ✓ 1100ms | ✓ 1611ms | ✓ 1352ms | http |
| 165.227.5.10:8888 | 否 | ✓ 978ms | ✓ 1216ms | ✓ 1280ms | ✓ 899ms | http |
| 183.249.5.117:22222 | ✓ 943ms | ✓ 1286ms | ✓ 857ms | ✓ 1388ms | 否 | http |
| 222.228.171.92:8080 | ✓ 1842ms | 否 | 否 | ✓ 1114ms | ✓ 1984ms | http |
| 113.59.32.162:22222 | ✓ 1161ms | ✓ 1438ms | ✓ 1205ms | 否 | 否 | http |
| 101.43.255.96:80 | ✓ 1047ms | ✓ 1479ms | ✓ 1116ms | ✓ 1464ms | ✓ 1087ms | http |
| 81.70.169.194:80 | ✓ 1103ms | ✓ 1462ms | ✓ 1352ms | ✓ 1422ms | ✓ 1158ms | http |
| 36.147.78.166:80 | 否 | ✓ 1768ms | ✓ 1886ms | ✓ 1802ms | ✓ 1812ms | http |
| 39.104.201.40:7890 | ✓ 1326ms | 否 | ✓ 1120ms | 否 | ✓ 1094ms | http |
| 103.104.99.29:80 | 否 | 否 | ✓ 1564ms | ✓ 1652ms | ✓ 1693ms | http |
| 103.104.99.89:80 | 否 | 否 | ✓ 1573ms | ✓ 1734ms | ✓ 1632ms | http |
| 91.238.104.171:2023 | ✓ 1449ms | ✓ 1530ms | ✓ 1293ms | ✓ 1807ms | ✓ 1591ms | http |
| 120.240.35.173:22222 | ✓ 1636ms | ✓ 1615ms | ✓ 1535ms | 否 | 否 | http |
| 183.249.5.214:22222 | ✓ 935ms | ✓ 1021ms | ✓ 878ms | 否 | 否 | http |
| 61.72.221.194:3128 | 否 | 否 | ✓ 1545ms | ✓ 1618ms | ✓ 1442ms | http |
| 116.236.189.91:29999 | ✓ 788ms | 否 | ✓ 836ms | ✓ 1052ms | 否 | http |
| 117.159.239.44:22222 | ✓ 896ms | ✓ 1254ms | ✓ 994ms | 否 | ✓ 1275ms | http |
| 113.59.32.141:22222 | ✓ 1131ms | ✓ 1469ms | 否 | ✓ 1354ms | 否 | http |
| 113.59.32.161:22222 | 否 | 否 | ✓ 1439ms | ✓ 1417ms | ✓ 1009ms | http |
| 35.234.17.221:8080 | ✓ 963ms | 否 | 否 | ✓ 1186ms | ✓ 975ms | http |
| 183.249.5.111:22222 | ✓ 912ms | ✓ 1288ms | ✓ 880ms | ✓ 1063ms | 否 | http |
| 125.129.39.95:3128 | ✓ 1913ms | ✓ 1934ms | ✓ 1391ms | ✓ 1364ms | ✓ 1245ms | http |
| 103.113.70.189:1081 | 否 | ✓ 938ms | 否 | ✓ 967ms | ✓ 814ms | http |
| 37.27.100.107:443 | ✓ 711ms | 否 | ✓ 1142ms | 否 | ✓ 1427ms | http |
| 37.27.100.80:443 | 否 | ✓ 1506ms | ✓ 549ms | ✓ 1540ms | ✓ 1683ms | http |
| 37.27.100.102:443 | 否 | ✓ 1434ms | ✓ 558ms | 否 | ✓ 1846ms | http |
| 37.27.100.79:443 | ✓ 708ms | ✓ 1686ms | ✓ 1455ms | 否 | 否 | http |
| 120.232.242.119:22222 | 否 | ✓ 1343ms | ✓ 1037ms | ✓ 1221ms | ✓ 1014ms | http |
| 115.231.181.40:8128 | ✓ 1002ms | ✓ 1218ms | ✓ 1022ms | ✓ 1431ms | ✓ 1053ms | http |
| 120.240.35.161:22222 | ✓ 1081ms | ✓ 1621ms | ✓ 1067ms | 否 | ✓ 1864ms | http |
| 45.125.67.37:8443 | ✓ 1686ms | 否 | ✓ 1709ms | 否 | ✓ 1924ms | http |
| 62.113.119.14:8080 | ✓ 597ms | 否 | ✓ 727ms | ✓ 1465ms | ✓ 1526ms | http |
| 185.213.20.105:3128 | ✓ 987ms | 否 | ✓ 1763ms | ✓ 1748ms | 否 | http |
| 120.238.159.189:22222 | ✓ 1028ms | ✓ 1343ms | ✓ 1164ms | ✓ 1248ms | ✓ 1001ms | http |
| 168.235.110.63:3128 | ✓ 887ms | 否 | ✓ 691ms | ✓ 1647ms | ✓ 1295ms | http |
| 47.101.149.27:9010 | ✓ 1483ms | 否 | ✓ 1465ms | ✓ 1419ms | 否 | http |
| 120.240.35.160:22222 | ✓ 1208ms | ✓ 1648ms | 否 | 否 | ✓ 1152ms | http |
| 222.184.48.235:22222 | ✓ 1014ms | ✓ 1279ms | ✓ 1053ms | ✓ 1353ms | ✓ 1041ms | http |
| 113.59.32.145:22222 | ✓ 1146ms | ✓ 1477ms | ✓ 1037ms | ✓ 1372ms | ✓ 1028ms | http |
| 120.240.35.177:22222 | ✓ 1158ms | ✓ 1637ms | ✓ 1243ms | ✓ 1467ms | ✓ 1182ms | http |
| 120.240.35.178:22222 | ✓ 1676ms | ✓ 1610ms | ✓ 1171ms | ✓ 1871ms | ✓ 1131ms | http |
| 121.230.8.39:1080 | ✓ 1420ms | ✓ 1784ms | ✓ 1470ms | ✓ 1692ms | ✓ 1305ms | http |
| 45.140.147.82:1081 | 否 | ✓ 1481ms | ✓ 1438ms | 否 | ✓ 1692ms | http |
| 91.233.223.147:3128 | ✓ 1555ms | 否 | 否 | ✓ 1947ms | ✓ 1944ms | http |
| 61.72.110.94:3128 | ✓ 1531ms | 否 | ✓ 862ms | 否 | ✓ 1107ms | http |
| 95.85.252.153:21064 | 否 | ✓ 1442ms | ✓ 877ms | ✓ 1648ms | ✓ 1495ms | http |
| 210.223.44.230:3128 | ✓ 1193ms | ✓ 1771ms | ✓ 1227ms | ✓ 1131ms | ✓ 889ms | http |
| 103.84.95.54:7890 | ✓ 794ms | 否 | ✓ 813ms | 否 | ✓ 871ms | http |
| 113.59.32.142:22222 | ✓ 1170ms | ✓ 1460ms | 否 | ✓ 1388ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1897ms | ✓ 1365ms | ✓ 1192ms | 否 | ✓ 1050ms | http |
| 45.140.147.155:1082 | ✓ 671ms | ✓ 1478ms | ✓ 1384ms | 否 | ✓ 1232ms | http |
| 103.236.64.247:8888 | ✓ 1574ms | ✓ 1675ms | ✓ 1849ms | 否 | 否 | http |
| 45.136.198.40:3128 | ✓ 1738ms | 否 | ✓ 1873ms | 否 | ✓ 1956ms | http |
| 103.82.23.118:5234 | 否 | 否 | ✓ 1965ms | ✓ 1946ms | ✓ 1545ms | http |
| 198.23.236.47:1111 | ✓ 821ms | ✓ 950ms | ✓ 1125ms | 否 | 否 | http |
| 222.184.48.252:22222 | ✓ 1014ms | ✓ 1266ms | ✓ 1018ms | ✓ 1321ms | ✓ 1070ms | http |
| 47.110.42.192:9003 | ✓ 1653ms | ✓ 1522ms | ✓ 1825ms | ✓ 1702ms | ✓ 1533ms | http |
| 104.238.30.40:59741 | ✓ 1728ms | 否 | ✓ 1839ms | 否 | ✓ 1999ms | http |
| 194.59.204.87:9080 | ✓ 470ms | ✓ 1445ms | ✓ 1104ms | 否 | 否 | http |
| 103.39.51.190:8080 | ✓ 1781ms | 否 | 否 | ✓ 1550ms | ✓ 1473ms | http |
| 47.77.180.205:1080 | ✓ 854ms | ✓ 948ms | ✓ 1030ms | ✓ 946ms | ✓ 715ms | http |
| 103.74.192.242:7890 | ✓ 1028ms | 否 | ✓ 1360ms | 否 | ✓ 1452ms | http |
| 121.40.231.103:7890 | 否 | 否 | ✓ 971ms | ✓ 1232ms | ✓ 1844ms | http |
| 222.184.48.242:22222 | ✓ 996ms | ✓ 1295ms | ✓ 1052ms | ✓ 1251ms | ✓ 935ms | http |
| 120.55.163.237:10086 | ✓ 984ms | ✓ 1170ms | ✓ 1053ms | ✓ 1210ms | ✓ 1335ms | http |
| 180.127.149.245:1080 | ✓ 1442ms | ✓ 1800ms | ✓ 1414ms | ✓ 1843ms | ✓ 1287ms | http |
| 106.14.205.114:483 | 否 | ✓ 1426ms | ✓ 1266ms | ✓ 1221ms | 否 | http |
| 14.56.177.44:3128 | ✓ 1302ms | ✓ 1533ms | ✓ 714ms | ✓ 1199ms | ✓ 929ms | http |
| 36.147.78.166:443 | ✓ 1834ms | 否 | ✓ 1774ms | 否 | ✓ 1550ms | http |
| 47.105.98.23:3128 | 否 | ✓ 1873ms | ✓ 1080ms | 否 | ✓ 1043ms | http |

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
