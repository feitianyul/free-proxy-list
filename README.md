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

最后更新：2026-04-05 14:01:51 UTC（2026-04-05 22:01:51 UTC+8）

**代理总数：84**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 84 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 84 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | ✓ 252ms | ✓ 1166ms | ✓ 545ms | ✓ 1164ms | ✓ 981ms | http |
| 147.161.210.140:8800 | ✓ 803ms | 否 | ✓ 1128ms | ✓ 993ms | ✓ 1201ms | http |
| 1.231.81.166:3128 | ✓ 1146ms | ✓ 1169ms | ✓ 1850ms | ✓ 1637ms | ✓ 1498ms | http |
| 159.223.71.162:443 | ✓ 1831ms | 否 | ✓ 1102ms | ✓ 1473ms | ✓ 1262ms | http |
| 159.223.71.162:8080 | ✓ 1829ms | 否 | ✓ 1079ms | ✓ 1305ms | ✓ 1447ms | http |
| 167.103.115.102:8800 | ✓ 1829ms | 否 | ✓ 1389ms | ✓ 1280ms | ✓ 1480ms | http |
| 113.160.132.26:8080 | ✓ 1742ms | 否 | ✓ 1377ms | ✓ 1512ms | ✓ 1680ms | http |
| 111.227.254.12:22222 | ✓ 1246ms | ✓ 1592ms | 否 | 否 | ✓ 1341ms | http |
| 185.41.152.110:3128 | ✓ 1920ms | ✓ 1869ms | ✓ 1761ms | 否 | ✓ 1591ms | http |
| 111.227.254.9:22222 | ✓ 1934ms | 否 | ✓ 1282ms | ✓ 1602ms | ✓ 1208ms | http |
| 38.145.203.34:8444 | ✓ 1242ms | ✓ 1574ms | ✓ 320ms | ✓ 1151ms | ✓ 1048ms | http |
| 167.103.34.108:8800 | ✓ 1137ms | 否 | 否 | ✓ 1373ms | ✓ 1336ms | http |
| 38.145.203.41:8453 | 否 | 否 | ✓ 1837ms | ✓ 1921ms | ✓ 1455ms | http |
| 167.103.144.127:8800 | ✓ 1555ms | 否 | ✓ 1631ms | 否 | ✓ 1701ms | http |
| 38.145.218.14:8446 | ✓ 882ms | ✓ 1913ms | ✓ 1230ms | ✓ 1452ms | ✓ 1648ms | http |
| 111.227.254.11:22222 | 否 | ✓ 1594ms | ✓ 1222ms | ✓ 1899ms | 否 | http |
| 45.167.125.21:999 | ✓ 808ms | ✓ 1827ms | ✓ 1372ms | ✓ 1702ms | ✓ 1428ms | http |
| 167.103.31.122:8800 | ✓ 1744ms | 否 | ✓ 1307ms | ✓ 1907ms | ✓ 1434ms | http |
| 163.44.126.97:3128 | ✓ 1692ms | 否 | 否 | ✓ 1988ms | ✓ 1929ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1197ms | ✓ 1461ms | ✓ 1171ms | http |
| 218.108.131.186:17890 | ✓ 1102ms | ✓ 1236ms | ✓ 1081ms | 否 | 否 | http |
| 183.232.248.73:7890 | ✓ 1186ms | ✓ 1345ms | ✓ 1119ms | 否 | 否 | http |
| 38.145.208.238:8451 | ✓ 1250ms | ✓ 1226ms | ✓ 1252ms | 否 | ✓ 1246ms | http |
| 147.161.239.240:8800 | ✓ 678ms | ✓ 1485ms | ✓ 1079ms | ✓ 1467ms | ✓ 1224ms | http |
| 34.101.184.164:3128 | ✓ 1681ms | 否 | ✓ 1166ms | ✓ 1446ms | ✓ 1211ms | http |
| 20.27.13.35:8561 | ✓ 1194ms | 否 | ✓ 809ms | ✓ 1019ms | ✓ 821ms | http |
| 210.223.44.230:3128 | ✓ 1724ms | ✓ 1922ms | ✓ 1163ms | ✓ 1487ms | ✓ 926ms | http |
| 115.231.181.40:8128 | ✓ 1032ms | ✓ 1373ms | ✓ 1079ms | ✓ 1470ms | ✓ 1244ms | http |
| 133.242.138.34:8100 | ✓ 938ms | ✓ 1665ms | ✓ 927ms | ✓ 1199ms | ✓ 1160ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1312ms | 否 | ✓ 1549ms | ✓ 1325ms | http |
| 209.38.154.7:1080 | ✓ 780ms | ✓ 980ms | ✓ 1364ms | 否 | ✓ 1992ms | http |
| 121.43.196.210:8222 | ✓ 1127ms | ✓ 1276ms | ✓ 1087ms | ✓ 1341ms | ✓ 1202ms | http |
| 20.78.118.91:8561 | ✓ 1773ms | ✓ 1903ms | ✓ 1269ms | ✓ 1373ms | ✓ 1024ms | http |
| 45.12.151.226:2829 | 否 | 否 | ✓ 1832ms | ✓ 1749ms | ✓ 1497ms | http |
| 38.34.179.53:8451 | ✓ 957ms | 否 | ✓ 567ms | ✓ 1079ms | ✓ 1264ms | http |
| 38.145.218.217:8450 | ✓ 558ms | ✓ 1989ms | ✓ 581ms | ✓ 980ms | ✓ 963ms | http |
| 38.145.218.212:8448 | ✓ 559ms | ✓ 1682ms | ✓ 671ms | ✓ 1036ms | ✓ 789ms | http |
| 38.34.179.74:8447 | ✓ 583ms | ✓ 994ms | ✓ 1241ms | ✓ 1618ms | ✓ 827ms | http |
| 38.34.179.75:8447 | ✓ 584ms | 否 | ✓ 1067ms | ✓ 1306ms | ✓ 739ms | http |
| 38.145.220.32:8449 | ✓ 1543ms | ✓ 1406ms | 否 | 否 | ✓ 918ms | http |
| 101.32.244.83:8080 | ✓ 1836ms | 否 | ✓ 1154ms | ✓ 1649ms | ✓ 1487ms | http |
| 120.92.212.16:7890 | ✓ 1267ms | ✓ 1637ms | 否 | ✓ 1730ms | ✓ 1383ms | http |
| 121.43.196.213:8222 | ✓ 1144ms | ✓ 1266ms | ✓ 1054ms | ✓ 1343ms | ✓ 1110ms | http |
| 114.55.226.123:10086 | ✓ 1248ms | 否 | ✓ 1285ms | ✓ 1519ms | ✓ 1239ms | http |
| 64.227.76.27:1080 | ✓ 392ms | ✓ 1704ms | ✓ 1595ms | ✓ 1685ms | ✓ 1179ms | http |
| 101.43.127.100:8877 | 否 | 否 | ✓ 1522ms | ✓ 1640ms | ✓ 1138ms | http |
| 38.34.179.69:8452 | ✓ 889ms | ✓ 1079ms | ✓ 884ms | ✓ 1162ms | ✓ 808ms | http |
| 38.145.203.32:8449 | ✓ 913ms | ✓ 926ms | ✓ 489ms | ✓ 1212ms | ✓ 1264ms | http |
| 38.145.220.198:8448 | ✓ 907ms | ✓ 1078ms | ✓ 1165ms | ✓ 1042ms | ✓ 841ms | http |
| 38.34.179.59:8444 | ✓ 901ms | ✓ 1982ms | ✓ 865ms | ✓ 1148ms | 否 | http |
| 38.145.218.101:8447 | ✓ 888ms | ✓ 1039ms | ✓ 597ms | ✓ 1586ms | ✓ 774ms | http |
| 38.145.203.98:8447 | ✓ 931ms | ✓ 1768ms | ✓ 782ms | ✓ 1978ms | ✓ 1309ms | http |
| 38.145.208.175:8447 | ✓ 908ms | 否 | ✓ 1575ms | ✓ 1344ms | ✓ 885ms | http |
| 114.237.77.244:1080 | 否 | ✓ 1472ms | ✓ 1125ms | ✓ 1414ms | ✓ 1256ms | http |
| 38.145.203.86:8450 | ✓ 904ms | ✓ 1464ms | ✓ 741ms | 否 | ✓ 1393ms | http |
| 111.227.254.10:22222 | ✓ 1202ms | ✓ 1638ms | 否 | ✓ 1523ms | ✓ 1338ms | http |
| 121.230.9.198:1080 | ✓ 1644ms | ✓ 1881ms | ✓ 1559ms | 否 | ✓ 1361ms | http |
| 38.34.179.73:8446 | ✓ 914ms | ✓ 1150ms | ✓ 1511ms | 否 | 否 | http |
| 217.217.249.160:8080 | ✓ 1988ms | 否 | ✓ 1045ms | 否 | ✓ 1741ms | http |
| 59.46.216.131:30001 | ✓ 1221ms | 否 | 否 | ✓ 1601ms | ✓ 1305ms | http |
| 20.210.39.153:8561 | ✓ 841ms | ✓ 1605ms | ✓ 688ms | ✓ 1003ms | ✓ 815ms | http |
| 20.78.26.206:8561 | ✓ 836ms | ✓ 1225ms | ✓ 1020ms | ✓ 1083ms | ✓ 798ms | http |
| 38.34.179.56:8450 | ✓ 1835ms | ✓ 1460ms | ✓ 1535ms | ✓ 1490ms | ✓ 939ms | http |
| 38.145.208.234:8446 | ✓ 627ms | ✓ 1961ms | ✓ 1545ms | 否 | 否 | http |
| 150.107.140.238:3128 | ✓ 1828ms | 否 | 否 | ✓ 1388ms | ✓ 1454ms | http |
| 86.53.183.16:1080 | ✓ 786ms | ✓ 1450ms | ✓ 1139ms | ✓ 1705ms | ✓ 1404ms | http |
| 38.145.218.210:8448 | ✓ 754ms | ✓ 873ms | ✓ 579ms | ✓ 1155ms | ✓ 1238ms | http |
| 181.205.39.236:8080 | ✓ 1972ms | 否 | ✓ 1270ms | ✓ 1714ms | ✓ 1464ms | http |
| 217.76.245.80:999 | ✓ 1056ms | 否 | ✓ 1317ms | ✓ 1535ms | ✓ 1242ms | http |
| 45.136.131.35:8445 | ✓ 1233ms | ✓ 1046ms | ✓ 1506ms | ✓ 1082ms | 否 | http |
| 45.136.131.29:8444 | ✓ 1218ms | ✓ 1434ms | ✓ 1126ms | ✓ 1686ms | 否 | http |
| 38.34.179.167:8451 | 否 | 否 | ✓ 366ms | ✓ 1034ms | ✓ 947ms | http |
| 38.34.179.173:8452 | 否 | 否 | ✓ 305ms | ✓ 1030ms | ✓ 1502ms | http |
| 212.58.132.5:8888 | ✓ 1154ms | 否 | ✓ 1451ms | 否 | ✓ 1299ms | http |
| 138.197.68.35:4857 | ✓ 681ms | ✓ 1966ms | 否 | 否 | ✓ 1561ms | http |
| 38.145.218.233:8450 | ✓ 323ms | ✓ 1395ms | ✓ 379ms | ✓ 988ms | ✓ 747ms | http |
| 38.34.179.63:8448 | ✓ 719ms | ✓ 1691ms | ✓ 352ms | ✓ 1189ms | ✓ 816ms | http |
| 38.145.220.27:8445 | ✓ 331ms | ✓ 978ms | ✓ 508ms | ✓ 1419ms | ✓ 1505ms | http |
| 5.104.87.17:8050 | 否 | 否 | ✓ 1303ms | ✓ 1294ms | ✓ 1217ms | http |
| 134.209.153.66:3128 | ✓ 1547ms | 否 | ✓ 1472ms | ✓ 1835ms | 否 | http |
| 5.104.87.17:8051 | 否 | 否 | ✓ 1665ms | ✓ 1479ms | ✓ 1529ms | http |
| 157.254.37.238:999 | ✓ 964ms | 否 | ✓ 1386ms | ✓ 1427ms | ✓ 1147ms | http |
| 45.136.130.166:8452 | ✓ 1853ms | 否 | ✓ 1328ms | ✓ 982ms | ✓ 910ms | http |
| 38.145.208.224:8445 | ✓ 1995ms | 否 | ✓ 1365ms | ✓ 903ms | ✓ 1157ms | http |

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
