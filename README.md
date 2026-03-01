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

最后更新：2026-03-01 20:32:00 UTC（2026-03-02 04:32:00 UTC+8）

**代理总数：77**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 1153ms | ✓ 1241ms | ✓ 1209ms | 否 | 否 | http |
| 141.11.210.35:1080 | ✓ 477ms | 否 | ✓ 970ms | ✓ 1360ms | ✓ 987ms | http |
| 190.9.109.196:999 | ✓ 1214ms | ✓ 1516ms | ✓ 1300ms | ✓ 1666ms | ✓ 1205ms | http |
| 121.128.121.54:3128 | ✓ 1296ms | ✓ 950ms | ✓ 709ms | 否 | ✓ 821ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1448ms | 否 | ✓ 1219ms | ✓ 847ms | http |
| 74.208.234.198:443 | 否 | ✓ 1565ms | 否 | ✓ 1485ms | ✓ 1112ms | http |
| 61.72.110.54:3128 | ✓ 1291ms | 否 | 否 | ✓ 1124ms | ✓ 1903ms | http |
| 59.46.216.131:30001 | ✓ 1139ms | ✓ 1455ms | ✓ 1252ms | ✓ 1358ms | ✓ 1173ms | http |
| 35.225.22.61:80 | 否 | ✓ 1187ms | ✓ 317ms | ✓ 1012ms | ✓ 813ms | http |
| 115.231.181.40:8128 | ✓ 920ms | ✓ 1277ms | ✓ 1063ms | ✓ 1208ms | ✓ 1021ms | http |
| 195.123.209.48:3128 | ✓ 1253ms | ✓ 1708ms | ✓ 1643ms | 否 | ✓ 1738ms | http |
| 14.56.107.244:3128 | 否 | ✓ 956ms | ✓ 1726ms | 否 | ✓ 1836ms | http |
| 120.92.212.16:7890 | ✓ 1078ms | 否 | ✓ 1560ms | ✓ 1316ms | ✓ 1042ms | http |
| 45.140.147.82:1081 | ✓ 1707ms | ✓ 1537ms | ✓ 1435ms | 否 | 否 | http |
| 45.125.67.37:8443 | 否 | 否 | ✓ 1266ms | ✓ 1505ms | ✓ 1270ms | http |
| 103.84.95.54:7890 | ✓ 990ms | 否 | ✓ 742ms | 否 | ✓ 719ms | http |
| 95.85.252.153:21064 | ✓ 914ms | ✓ 1883ms | ✓ 1445ms | 否 | 否 | http |
| 180.127.149.228:1080 | ✓ 1003ms | ✓ 1319ms | ✓ 1062ms | ✓ 1234ms | ✓ 1028ms | http |
| 39.104.201.40:7890 | ✓ 1019ms | ✓ 1288ms | ✓ 1095ms | ✓ 1295ms | ✓ 1060ms | http |
| 101.43.255.96:80 | ✓ 1061ms | ✓ 1485ms | ✓ 1058ms | ✓ 1304ms | ✓ 1122ms | http |
| 81.70.169.194:80 | ✓ 1101ms | ✓ 1508ms | ✓ 1156ms | ✓ 1281ms | ✓ 1045ms | http |
| 36.147.78.166:80 | ✓ 1793ms | 否 | ✓ 1743ms | ✓ 1983ms | 否 | http |
| 14.56.177.44:3128 | ✓ 1098ms | 否 | ✓ 876ms | ✓ 1101ms | ✓ 914ms | http |
| 36.147.78.166:443 | ✓ 1828ms | ✓ 1832ms | ✓ 1825ms | ✓ 1716ms | 否 | http |
| 222.228.171.92:8080 | 否 | 否 | ✓ 1266ms | ✓ 1355ms | ✓ 945ms | http |
| 35.234.17.221:8080 | ✓ 1005ms | ✓ 1803ms | 否 | ✓ 1650ms | 否 | http |
| 209.38.183.26:3128 | ✓ 494ms | ✓ 1666ms | ✓ 1660ms | ✓ 1660ms | ✓ 1323ms | http |
| 47.105.98.23:3128 | 否 | ✓ 1466ms | 否 | ✓ 1345ms | ✓ 1003ms | http |
| 103.191.169.130:1111 | 否 | 否 | ✓ 1909ms | ✓ 1706ms | ✓ 1627ms | http |
| 61.72.221.194:3128 | 否 | 否 | ✓ 1254ms | ✓ 1631ms | ✓ 1430ms | http |
| 120.92.212.16:8890 | ✓ 1375ms | ✓ 1306ms | 否 | 否 | ✓ 1815ms | http |
| 47.77.180.205:1080 | ✓ 1283ms | ✓ 1042ms | ✓ 638ms | ✓ 969ms | ✓ 893ms | http |
| 162.240.154.26:3128 | ✓ 1384ms | ✓ 1094ms | ✓ 898ms | ✓ 1061ms | ✓ 1044ms | http |
| 167.160.184.231:6005 | ✓ 590ms | ✓ 1256ms | ✓ 1160ms | ✓ 1227ms | ✓ 994ms | http |
| 2.56.178.131:443 | ✓ 1070ms | 否 | ✓ 859ms | 否 | ✓ 1681ms | http |
| 111.79.111.126:3128 | 否 | ✓ 1768ms | 否 | ✓ 1887ms | ✓ 1825ms | http |
| 45.140.147.82:1082 | ✓ 516ms | ✓ 1907ms | ✓ 1657ms | 否 | 否 | http |
| 61.72.110.94:3128 | ✓ 1997ms | ✓ 1030ms | ✓ 1026ms | 否 | 否 | http |
| 85.198.84.77:10808 | 否 | 否 | ✓ 1837ms | ✓ 1987ms | ✓ 1711ms | http |
| 91.238.104.172:2024 | ✓ 767ms | ✓ 1891ms | ✓ 1988ms | 否 | 否 | http |
| 125.128.12.84:3128 | 否 | 否 | ✓ 1552ms | ✓ 1826ms | ✓ 1325ms | http |
| 180.127.149.225:1080 | ✓ 1077ms | ✓ 1313ms | ✓ 999ms | ✓ 1294ms | ✓ 1054ms | http |
| 121.230.8.211:1080 | 否 | ✓ 1461ms | ✓ 1103ms | ✓ 1510ms | ✓ 1126ms | http |
| 210.223.44.230:3128 | ✓ 805ms | ✓ 1413ms | ✓ 968ms | ✓ 1348ms | ✓ 1116ms | http |
| 168.235.110.63:3128 | ✓ 1305ms | 否 | ✓ 1433ms | ✓ 1674ms | ✓ 1041ms | http |
| 62.113.119.14:8080 | ✓ 1162ms | ✓ 1601ms | ✓ 1003ms | ✓ 1526ms | ✓ 1141ms | http |
| 37.27.100.102:443 | ✓ 1430ms | 否 | ✓ 940ms | 否 | ✓ 1769ms | http |
| 34.101.184.164:3128 | ✓ 1668ms | 否 | ✓ 965ms | ✓ 1515ms | ✓ 1262ms | http |
| 125.128.12.194:3128 | ✓ 1437ms | ✓ 1634ms | ✓ 1780ms | 否 | 否 | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1227ms | ✓ 1171ms | ✓ 1607ms | http |
| 165.227.5.10:8888 | 否 | ✓ 1142ms | ✓ 675ms | ✓ 880ms | ✓ 1791ms | http |
| 94.177.131.12:3128 | ✓ 1576ms | 否 | ✓ 698ms | ✓ 882ms | ✓ 733ms | http |
| 121.40.231.103:7890 | ✓ 1469ms | ✓ 1761ms | 否 | 否 | ✓ 1854ms | http |
| 37.27.100.112:443 | 否 | ✓ 1589ms | ✓ 1652ms | 否 | ✓ 1226ms | http |
| 91.238.104.171:2023 | ✓ 1021ms | ✓ 1889ms | ✓ 1586ms | 否 | 否 | http |
| 45.136.198.40:3128 | ✓ 1332ms | ✓ 1959ms | ✓ 1571ms | 否 | ✓ 1824ms | http |
| 38.180.2.107:3128 | ✓ 1334ms | ✓ 1951ms | ✓ 1574ms | 否 | ✓ 1923ms | http |
| 45.129.141.143:3128 | ✓ 1315ms | ✓ 1841ms | ✓ 1952ms | 否 | 否 | http |
| 194.59.204.87:9080 | ✓ 532ms | ✓ 1587ms | ✓ 1190ms | 否 | 否 | http |
| 124.121.2.247:8080 | 否 | 否 | ✓ 1744ms | ✓ 1692ms | ✓ 1606ms | http |
| 103.39.51.190:8080 | ✓ 1861ms | 否 | ✓ 1687ms | ✓ 1548ms | ✓ 1554ms | http |
| 217.77.102.18:3128 | ✓ 1764ms | 否 | ✓ 1745ms | 否 | ✓ 1864ms | http |
| 37.27.100.80:443 | ✓ 1864ms | ✓ 1588ms | ✓ 810ms | 否 | 否 | http |
| 147.45.60.34:1082 | 否 | 否 | ✓ 788ms | ✓ 1547ms | ✓ 877ms | http |
| 201.150.116.32:999 | ✓ 1764ms | ✓ 1318ms | 否 | 否 | ✓ 1211ms | http |
| 113.176.92.71:3128 | 否 | ✓ 1611ms | 否 | ✓ 1931ms | ✓ 1947ms | http |
| 103.82.23.118:5216 | ✓ 1684ms | 否 | ✓ 1465ms | 否 | ✓ 1785ms | http |
| 103.236.64.247:8888 | ✓ 1812ms | ✓ 1390ms | 否 | ✓ 1684ms | 否 | http |
| 142.171.85.32:1080 | ✓ 1363ms | ✓ 1259ms | ✓ 994ms | ✓ 970ms | ✓ 1108ms | http |
| 106.14.205.114:483 | ✓ 1099ms | ✓ 1097ms | ✓ 1267ms | ✓ 1188ms | ✓ 939ms | http |
| 120.132.97.88:7897 | ✓ 1020ms | ✓ 1297ms | ✓ 1009ms | ✓ 1326ms | ✓ 943ms | http |
| 198.23.236.47:1111 | 否 | 否 | ✓ 1733ms | ✓ 1333ms | ✓ 1028ms | http |
| 103.74.192.242:7890 | ✓ 932ms | ✓ 1469ms | ✓ 1230ms | ✓ 1071ms | ✓ 790ms | http |
| 199.68.217.2:3128 | ✓ 752ms | ✓ 1274ms | ✓ 1298ms | ✓ 1099ms | ✓ 752ms | http |
| 120.55.163.237:10086 | ✓ 957ms | ✓ 1148ms | ✓ 979ms | ✓ 1197ms | ✓ 996ms | http |
| 47.110.42.192:9003 | ✓ 1744ms | ✓ 1742ms | ✓ 1660ms | ✓ 1721ms | ✓ 1682ms | http |
| 83.229.73.113:13554 | ✓ 910ms | ✓ 1904ms | 否 | 否 | ✓ 1749ms | http |

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
