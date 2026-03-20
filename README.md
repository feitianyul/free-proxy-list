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

最后更新：2026-03-20 23:24:38 UTC（2026-03-21 07:24:38 UTC+8）

**代理总数：93**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 92 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 93 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1765ms | ✓ 1413ms | ✓ 1394ms | ✓ 1191ms | ✓ 1231ms | http |
| 137.220.150.152:6005 | ✓ 916ms | 否 | ✓ 940ms | ✓ 1319ms | ✓ 1082ms | http |
| 219.117.204.211:7799 | ✓ 1766ms | ✓ 1646ms | ✓ 1115ms | 否 | ✓ 839ms | http |
| 137.220.150.170:6005 | ✓ 1293ms | 否 | ✓ 921ms | ✓ 1427ms | ✓ 1391ms | http |
| 167.103.34.108:8800 | ✓ 1474ms | 否 | ✓ 1360ms | ✓ 1680ms | 否 | http |
| 45.167.124.52:8080 | ✓ 1229ms | ✓ 1898ms | ✓ 513ms | ✓ 1534ms | ✓ 1274ms | http |
| 45.136.130.196:8448 | 否 | ✓ 914ms | ✓ 814ms | ✓ 918ms | ✓ 796ms | http |
| 38.34.179.105:8449 | 否 | ✓ 925ms | ✓ 824ms | ✓ 1224ms | ✓ 729ms | http |
| 183.249.5.117:22222 | 否 | 否 | ✓ 889ms | ✓ 1207ms | ✓ 1076ms | http |
| 43.99.54.236:5555 | ✓ 882ms | ✓ 1150ms | ✓ 875ms | ✓ 1038ms | ✓ 828ms | http |
| 93.183.69.162:1080 | ✓ 585ms | 否 | ✓ 1386ms | 否 | ✓ 1668ms | http |
| 113.160.132.26:8080 | ✓ 1514ms | ✓ 1444ms | ✓ 1705ms | ✓ 1347ms | ✓ 1081ms | http |
| 194.67.99.223:1080 | ✓ 1353ms | 否 | 否 | ✓ 1659ms | ✓ 1547ms | http |
| 174.138.24.77:1080 | ✓ 1369ms | 否 | 否 | ✓ 1911ms | ✓ 1070ms | http |
| 38.34.179.83:8448 | 否 | ✓ 1432ms | ✓ 955ms | 否 | ✓ 1024ms | http |
| 167.103.31.122:8800 | ✓ 1752ms | 否 | ✓ 1572ms | ✓ 1916ms | 否 | http |
| 137.184.1.87:3128 | ✓ 931ms | ✓ 1067ms | ✓ 1391ms | ✓ 988ms | ✓ 740ms | http |
| 103.82.23.118:5171 | ✓ 1650ms | 否 | ✓ 1559ms | ✓ 1928ms | ✓ 1650ms | http |
| 120.92.212.16:7890 | ✓ 1204ms | ✓ 1689ms | 否 | 否 | ✓ 1157ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1095ms | ✓ 1455ms | ✓ 1137ms | http |
| 45.136.131.58:8450 | ✓ 1741ms | 否 | ✓ 766ms | ✓ 1243ms | 否 | http |
| 38.34.179.96:8451 | ✓ 738ms | ✓ 1028ms | 否 | ✓ 1193ms | ✓ 704ms | http |
| 35.225.22.61:80 | ✓ 779ms | ✓ 1240ms | ✓ 917ms | ✓ 1170ms | 否 | http |
| 45.136.131.35:8452 | ✓ 778ms | ✓ 1009ms | ✓ 548ms | ✓ 963ms | ✓ 841ms | http |
| 147.161.239.240:8800 | ✓ 1083ms | ✓ 1462ms | ✓ 1421ms | ✓ 1525ms | ✓ 1280ms | http |
| 101.43.127.100:8877 | ✓ 1085ms | ✓ 1274ms | ✓ 1025ms | ✓ 1241ms | ✓ 1044ms | http |
| 91.238.105.64:2024 | ✓ 1114ms | ✓ 1538ms | 否 | 否 | ✓ 1552ms | http |
| 137.220.150.22:6005 | ✓ 1645ms | 否 | ✓ 1360ms | ✓ 1351ms | ✓ 1057ms | http |
| 137.220.151.110:6005 | ✓ 1637ms | 否 | ✓ 1260ms | ✓ 1687ms | ✓ 1069ms | http |
| 217.174.244.117:3129 | 否 | ✓ 1683ms | 否 | ✓ 1981ms | ✓ 1718ms | http |
| 38.34.178.155:8448 | ✓ 993ms | ✓ 1206ms | 否 | ✓ 1567ms | 否 | http |
| 38.34.179.63:8448 | ✓ 1618ms | ✓ 1109ms | ✓ 753ms | ✓ 948ms | ✓ 745ms | http |
| 59.46.216.131:30001 | ✓ 1151ms | ✓ 1633ms | ✓ 1272ms | ✓ 1542ms | ✓ 1283ms | http |
| 183.249.5.111:22222 | ✓ 878ms | ✓ 1251ms | ✓ 1027ms | ✓ 1200ms | ✓ 1006ms | http |
| 111.201.98.211:7890 | ✓ 1035ms | ✓ 1322ms | ✓ 1020ms | ✓ 1386ms | ✓ 1074ms | http |
| 183.249.5.110:22222 | ✓ 999ms | ✓ 1212ms | ✓ 974ms | ✓ 1266ms | ✓ 1476ms | http |
| 83.219.250.8:62920 | ✓ 1045ms | ✓ 1440ms | ✓ 1528ms | 否 | ✓ 1636ms | http |
| 1.225.116.115:1080 | ✓ 1842ms | ✓ 1737ms | ✓ 1537ms | ✓ 1280ms | ✓ 1221ms | http |
| 194.147.115.50:3128 | ✓ 810ms | ✓ 1972ms | ✓ 950ms | 否 | 否 | http |
| 38.145.208.244:8448 | ✓ 835ms | ✓ 887ms | ✓ 794ms | ✓ 1099ms | ✓ 766ms | http |
| 38.34.183.222:8453 | ✓ 842ms | ✓ 842ms | ✓ 589ms | ✓ 1260ms | ✓ 1048ms | http |
| 38.34.179.60:8450 | ✓ 1097ms | 否 | ✓ 411ms | ✓ 1441ms | 否 | http |
| 72.56.79.129:1080 | ✓ 1162ms | 否 | 否 | ✓ 1869ms | ✓ 1448ms | http |
| 14.225.212.37:7890 | ✓ 1495ms | ✓ 1981ms | ✓ 1102ms | ✓ 1485ms | 否 | http |
| 222.184.48.251:22222 | ✓ 1141ms | ✓ 1800ms | ✓ 1470ms | 否 | 否 | http |
| 137.220.150.104:6005 | 否 | 否 | ✓ 1975ms | ✓ 1643ms | ✓ 1474ms | http |
| 38.95.77.16:6005 | ✓ 844ms | ✓ 1498ms | ✓ 1108ms | ✓ 1398ms | ✓ 1481ms | http |
| 38.34.179.14:8450 | ✓ 861ms | 否 | ✓ 1646ms | ✓ 1968ms | ✓ 871ms | http |
| 172.212.68.37:3128 | ✓ 191ms | ✓ 1622ms | ✓ 732ms | ✓ 1190ms | ✓ 1013ms | http |
| 38.145.220.39:8453 | ✓ 598ms | ✓ 1078ms | ✓ 995ms | ✓ 1166ms | ✓ 1896ms | http |
| 133.242.138.34:8100 | ✓ 1623ms | 否 | ✓ 1464ms | ✓ 1975ms | ✓ 1000ms | http |
| 106.75.15.167:7890 | ✓ 1403ms | ✓ 1719ms | 否 | 否 | ✓ 1091ms | http |
| 190.12.150.244:999 | ✓ 1246ms | ✓ 1780ms | ✓ 888ms | ✓ 1772ms | ✓ 1465ms | http |
| 183.249.5.105:22222 | ✓ 1090ms | ✓ 1352ms | ✓ 947ms | ✓ 1171ms | ✓ 935ms | http |
| 181.78.194.249:999 | ✓ 1497ms | ✓ 1853ms | ✓ 744ms | ✓ 1642ms | ✓ 1309ms | http |
| 15.204.151.141:3128 | 否 | 否 | ✓ 1815ms | ✓ 1499ms | ✓ 1230ms | http |
| 142.171.224.229:7890 | ✓ 846ms | ✓ 974ms | ✓ 825ms | ✓ 935ms | ✓ 721ms | http |
| 88.80.150.82:8080 | ✓ 1399ms | 否 | 否 | ✓ 1817ms | ✓ 1601ms | https |
| 144.31.52.77:3128 | ✓ 1057ms | ✓ 1945ms | ✓ 551ms | ✓ 1930ms | ✓ 1321ms | http |
| 115.231.181.40:8128 | ✓ 1347ms | ✓ 1370ms | 否 | ✓ 1450ms | ✓ 1119ms | http |
| 103.113.70.189:1081 | ✓ 265ms | ✓ 1881ms | ✓ 684ms | ✓ 931ms | ✓ 846ms | http |
| 45.136.131.49:8447 | ✓ 748ms | ✓ 1389ms | ✓ 288ms | ✓ 1158ms | ✓ 708ms | http |
| 77.110.113.24:40000 | ✓ 1045ms | 否 | ✓ 507ms | ✓ 1748ms | ✓ 1696ms | http |
| 45.136.130.186:8451 | ✓ 1013ms | ✓ 841ms | ✓ 394ms | ✓ 1357ms | ✓ 1802ms | http |
| 210.45.70.16:7895 | ✓ 1255ms | ✓ 1554ms | 否 | ✓ 1617ms | ✓ 1469ms | http |
| 106.117.208.101:7890 | ✓ 1149ms | ✓ 1480ms | ✓ 1238ms | ✓ 1470ms | ✓ 1153ms | http |
| 147.45.67.148:8080 | ✓ 976ms | ✓ 1894ms | ✓ 1286ms | 否 | 否 | http |
| 38.34.179.150:8449 | ✓ 681ms | ✓ 973ms | ✓ 663ms | ✓ 1054ms | ✓ 1477ms | http |
| 103.183.10.169:3125 | 否 | 否 | ✓ 1871ms | ✓ 1737ms | ✓ 1707ms | http |
| 223.16.170.103:80 | 否 | ✓ 1916ms | ✓ 1453ms | ✓ 1746ms | ✓ 1335ms | http |
| 45.129.141.143:3128 | ✓ 1126ms | ✓ 1667ms | ✓ 1813ms | ✓ 1730ms | ✓ 1637ms | http |
| 62.234.206.73:3128 | ✓ 1084ms | ✓ 1415ms | ✓ 1047ms | ✓ 1416ms | ✓ 1170ms | http |
| 45.136.198.40:3128 | ✓ 1146ms | ✓ 1574ms | ✓ 1688ms | 否 | ✓ 1767ms | http |
| 139.99.238.95:8080 | ✓ 1540ms | 否 | ✓ 1152ms | 否 | ✓ 1171ms | http |
| 38.145.220.33:8448 | 否 | ✓ 1631ms | ✓ 1313ms | 否 | ✓ 1006ms | http |
| 38.145.208.165:8450 | ✓ 344ms | ✓ 961ms | ✓ 321ms | ✓ 924ms | ✓ 720ms | http |
| 38.34.179.229:8445 | ✓ 968ms | ✓ 897ms | ✓ 580ms | ✓ 1675ms | ✓ 1439ms | http |
| 38.145.208.239:8445 | ✓ 342ms | ✓ 1119ms | ✓ 1202ms | ✓ 1411ms | ✓ 866ms | http |
| 180.103.19.166:1080 | ✓ 1480ms | ✓ 1481ms | ✓ 1323ms | ✓ 1796ms | ✓ 1391ms | http |
| 193.233.84.129:1080 | 否 | 否 | ✓ 1790ms | ✓ 1865ms | ✓ 1415ms | http |
| 222.184.48.252:22222 | ✓ 1341ms | ✓ 1305ms | ✓ 1059ms | 否 | ✓ 1556ms | http |
| 125.76.214.178:8091 | ✓ 1131ms | ✓ 1626ms | ✓ 1216ms | ✓ 1677ms | ✓ 1221ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1656ms | ✓ 1605ms | ✓ 1610ms | http |
| 38.34.179.162:8451 | ✓ 1213ms | 否 | ✓ 790ms | 否 | ✓ 1644ms | http |
| 121.230.9.252:1080 | ✓ 1403ms | ✓ 1789ms | ✓ 1342ms | ✓ 1657ms | ✓ 1132ms | http |
| 38.34.179.67:8443 | ✓ 845ms | ✓ 887ms | ✓ 1019ms | ✓ 965ms | ✓ 745ms | http |
| 38.34.179.66:8443 | ✓ 867ms | ✓ 902ms | ✓ 1007ms | ✓ 950ms | ✓ 764ms | http |
| 38.34.179.64:8443 | ✓ 345ms | ✓ 952ms | ✓ 294ms | ✓ 915ms | ✓ 807ms | http |
| 38.145.220.11:8445 | 否 | ✓ 981ms | ✓ 999ms | ✓ 1338ms | ✓ 738ms | http |
| 38.34.183.225:8450 | ✓ 864ms | 否 | ✓ 938ms | ✓ 1201ms | ✓ 1649ms | http |
| 38.34.179.190:8450 | ✓ 889ms | ✓ 917ms | ✓ 496ms | ✓ 1489ms | ✓ 1420ms | http |
| 38.34.179.213:8450 | ✓ 1861ms | ✓ 1029ms | 否 | ✓ 1815ms | 否 | http |
| 103.16.118.78:8070 | ✓ 1608ms | 否 | ✓ 1606ms | ✓ 1945ms | 否 | http |

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
