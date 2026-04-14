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

最后更新：2026-04-14 08:45:50 UTC（2026-04-14 16:45:50 UTC+8）

**代理总数：76**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 34.71.229.255:3128 | ✓ 713ms | 否 | ✓ 1300ms | ✓ 1000ms | ✓ 1011ms | http |
| 46.101.126.84:8888 | ✓ 391ms | 否 | ✓ 1400ms | ✓ 1603ms | ✓ 1353ms | http |
| 147.161.210.140:8800 | ✓ 1737ms | ✓ 1369ms | ✓ 847ms | ✓ 1187ms | ✓ 1031ms | http |
| 167.103.115.102:8800 | ✓ 1363ms | 否 | ✓ 1140ms | ✓ 1341ms | ✓ 1253ms | http |
| 43.156.132.113:3128 | ✓ 930ms | 否 | ✓ 1452ms | ✓ 1613ms | ✓ 1289ms | http |
| 129.154.48.5:1080 | ✓ 1762ms | ✓ 1531ms | ✓ 1860ms | ✓ 1421ms | ✓ 1187ms | http |
| 113.160.132.26:8080 | ✓ 1689ms | 否 | 否 | ✓ 1790ms | ✓ 1119ms | http |
| 78.11.96.22:8888 | ✓ 824ms | ✓ 1863ms | ✓ 1135ms | ✓ 1816ms | ✓ 1252ms | http |
| 35.225.22.61:80 | ✓ 405ms | ✓ 1107ms | 否 | 否 | ✓ 954ms | http |
| 157.230.178.216:8088 | ✓ 1240ms | ✓ 1834ms | ✓ 134ms | ✓ 989ms | ✓ 727ms | http |
| 167.103.34.108:8800 | ✓ 1735ms | 否 | ✓ 1367ms | ✓ 1602ms | ✓ 1306ms | http |
| 85.239.59.252:7890 | ✓ 1001ms | 否 | ✓ 1022ms | 否 | ✓ 1716ms | http |
| 159.223.225.118:8888 | ✓ 894ms | 否 | ✓ 1128ms | 否 | ✓ 1572ms | http |
| 45.167.125.21:999 | ✓ 1330ms | 否 | ✓ 1275ms | 否 | ✓ 1476ms | http |
| 147.161.239.240:8800 | ✓ 609ms | 否 | ✓ 1348ms | ✓ 1558ms | ✓ 1202ms | http |
| 167.103.144.127:8800 | ✓ 1575ms | 否 | ✓ 1116ms | ✓ 1464ms | ✓ 1379ms | http |
| 167.103.31.122:8800 | ✓ 1572ms | 否 | ✓ 1490ms | ✓ 1918ms | ✓ 1492ms | http |
| 45.149.92.147:5001 | ✓ 981ms | 否 | ✓ 1050ms | ✓ 1222ms | 否 | http |
| 108.131.109.106:3129 | ✓ 733ms | 否 | ✓ 947ms | ✓ 1983ms | ✓ 1416ms | http |
| 52.59.218.12:3129 | ✓ 692ms | 否 | ✓ 940ms | 否 | ✓ 1519ms | http |
| 18.170.25.193:3129 | ✓ 830ms | 否 | 否 | ✓ 1731ms | ✓ 1325ms | http |
| 192.71.213.85:9812 | ✓ 863ms | 否 | ✓ 1446ms | ✓ 1915ms | 否 | http |
| 192.71.213.85:5555 | ✓ 862ms | 否 | ✓ 1449ms | ✓ 1920ms | 否 | http |
| 15.160.132.166:3129 | ✓ 644ms | 否 | 否 | ✓ 1633ms | ✓ 1790ms | http |
| 184.72.0.186:3129 | ✓ 1547ms | 否 | ✓ 1895ms | 否 | ✓ 1939ms | http |
| 101.43.127.100:8877 | 否 | 否 | ✓ 1952ms | ✓ 1407ms | ✓ 1072ms | http |
| 59.46.216.131:30001 | ✓ 1204ms | ✓ 1563ms | ✓ 1373ms | 否 | 否 | http |
| 194.87.85.207:1080 | ✓ 1648ms | 否 | ✓ 1808ms | ✓ 1871ms | ✓ 1855ms | http |
| 177.234.217.88:999 | ✓ 1101ms | ✓ 1692ms | ✓ 1075ms | ✓ 1732ms | ✓ 1458ms | http |
| 103.113.70.189:1081 | ✓ 536ms | ✓ 1817ms | ✓ 976ms | 否 | ✓ 1458ms | http |
| 16.62.123.236:3128 | ✓ 970ms | 否 | ✓ 1085ms | ✓ 1972ms | ✓ 1604ms | http |
| 38.180.2.107:3128 | ✓ 848ms | ✓ 1529ms | ✓ 1807ms | 否 | ✓ 1717ms | http |
| 157.0.142.246:10061 | ✓ 1151ms | ✓ 1480ms | ✓ 1220ms | ✓ 1511ms | ✓ 1187ms | http |
| 13.53.139.178:3128 | ✓ 1541ms | 否 | ✓ 1656ms | ✓ 1854ms | ✓ 1480ms | http |
| 218.108.131.186:17890 | ✓ 1034ms | ✓ 1718ms | ✓ 1349ms | ✓ 1971ms | ✓ 1146ms | http |
| 103.113.70.189:1082 | ✓ 762ms | 否 | ✓ 915ms | 否 | ✓ 1426ms | http |
| 36.141.21.200:7890 | 否 | 否 | ✓ 1145ms | ✓ 1858ms | ✓ 1315ms | http |
| 62.113.119.14:8080 | ✓ 540ms | 否 | ✓ 809ms | ✓ 1465ms | ✓ 1133ms | http |
| 138.124.99.216:8888 | ✓ 523ms | ✓ 1523ms | ✓ 1529ms | ✓ 1643ms | ✓ 1594ms | http |
| 144.31.27.49:1080 | ✓ 1076ms | ✓ 1863ms | ✓ 1307ms | 否 | ✓ 1706ms | http |
| 212.58.132.5:8888 | ✓ 1209ms | 否 | ✓ 1734ms | ✓ 1480ms | ✓ 1260ms | http |
| 210.77.22.250:7890 | ✓ 1121ms | ✓ 1473ms | ✓ 1341ms | ✓ 1452ms | ✓ 1313ms | http |
| 34.246.183.20:8081 | ✓ 1689ms | 否 | ✓ 1724ms | 否 | ✓ 1456ms | http |
| 2.27.18.184:1080 | ✓ 1699ms | ✓ 1808ms | 否 | 否 | ✓ 1300ms | http |
| 72.56.84.21:1080 | ✓ 967ms | 否 | ✓ 1746ms | 否 | ✓ 1701ms | http |
| 5.104.87.17:8051 | 否 | 否 | ✓ 1315ms | ✓ 1708ms | ✓ 1730ms | http |
| 104.129.202.127:12354 | 否 | 否 | ✓ 1375ms | ✓ 1043ms | ✓ 1000ms | http |
| 207.254.71.62:8088 | ✓ 1356ms | ✓ 1635ms | 否 | ✓ 1880ms | ✓ 1442ms | http |
| 3.71.175.73:3129 | ✓ 1075ms | 否 | ✓ 1405ms | 否 | ✓ 1306ms | http |
| 192.71.213.85:9091 | ✓ 967ms | 否 | ✓ 1250ms | ✓ 1865ms | 否 | http |
| 51.48.136.148:3129 | ✓ 1117ms | 否 | ✓ 1673ms | 否 | ✓ 1829ms | http |
| 171.244.130.36:3128 | ✓ 1650ms | 否 | ✓ 1895ms | ✓ 1986ms | ✓ 1951ms | http |
| 112.78.181.94:3128 | 否 | 否 | ✓ 1634ms | ✓ 1762ms | ✓ 1724ms | http |
| 51.17.154.141:3128 | ✓ 1305ms | 否 | ✓ 971ms | 否 | ✓ 1490ms | http |
| 193.181.35.210:8118 | ✓ 698ms | 否 | ✓ 989ms | 否 | ✓ 1544ms | http |
| 45.140.147.155:1082 | ✓ 586ms | ✓ 1318ms | ✓ 709ms | ✓ 1349ms | 否 | http |
| 104.129.202.127:10810 | 否 | ✓ 1724ms | ✓ 938ms | ✓ 1038ms | ✓ 887ms | http |
| 5.255.123.43:1080 | ✓ 631ms | ✓ 1839ms | 否 | 否 | ✓ 1487ms | http |
| 108.131.109.106:41392 | ✓ 548ms | 否 | ✓ 1480ms | ✓ 1423ms | ✓ 1534ms | http |
| 3.96.208.241:22894 | ✓ 1243ms | 否 | ✓ 1100ms | ✓ 1658ms | ✓ 1592ms | http |
| 52.16.215.4:30568 | ✓ 566ms | 否 | ✓ 1477ms | ✓ 1972ms | ✓ 1112ms | http |
| 52.47.115.41:24427 | ✓ 621ms | 否 | ✓ 1680ms | ✓ 1583ms | ✓ 1586ms | http |
| 35.178.250.178:443 | ✓ 1419ms | 否 | ✓ 1942ms | ✓ 1425ms | ✓ 1407ms | http |
| 47.84.131.156:8100 | ✓ 1298ms | ✓ 1950ms | 否 | 否 | ✓ 1338ms | http |
| 54.241.232.209:3128 | ✓ 1153ms | 否 | ✓ 1334ms | 否 | ✓ 1876ms | http |
| 52.59.51.29:45726 | ✓ 1582ms | 否 | ✓ 1640ms | 否 | ✓ 1561ms | http |
| 210.223.44.230:3128 | ✓ 1558ms | ✓ 1497ms | ✓ 1094ms | 否 | ✓ 1817ms | http |
| 2.27.8.54:1080 | ✓ 1023ms | ✓ 1532ms | ✓ 893ms | ✓ 1597ms | ✓ 960ms | http |
| 204.152.223.225:9050 | ✓ 254ms | 否 | 否 | ✓ 1990ms | ✓ 749ms | http |
| 107.172.102.234:40621 | ✓ 1119ms | 否 | ✓ 1067ms | ✓ 1082ms | 否 | http |
| 34.101.184.164:3128 | ✓ 1041ms | 否 | ✓ 1138ms | ✓ 1503ms | ✓ 1599ms | http |
| 101.32.243.189:80 | ✓ 1396ms | 否 | 否 | ✓ 1942ms | ✓ 1527ms | http |
| 61.52.131.172:8443 | ✓ 1110ms | ✓ 1388ms | ✓ 1182ms | ✓ 1395ms | ✓ 1162ms | http |
| 204.157.251.178:999 | ✓ 1690ms | ✓ 1373ms | ✓ 1698ms | 否 | ✓ 1248ms | http |
| 223.16.170.103:3128 | ✓ 1413ms | 否 | ✓ 1339ms | ✓ 1386ms | ✓ 1363ms | http |
| 103.39.51.207:8080 | ✓ 1752ms | 否 | ✓ 1741ms | ✓ 1904ms | ✓ 1829ms | http |

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
