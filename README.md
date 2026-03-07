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

最后更新：2026-03-07 08:36:43 UTC（2026-03-07 16:36:43 UTC+8）

**代理总数：88**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 88 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 88 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 1660ms | ✓ 1196ms | ✓ 1295ms | ✓ 878ms | ✓ 810ms | http |
| 205.209.118.30:3138 | ✓ 1012ms | 否 | ✓ 1424ms | ✓ 1416ms | ✓ 1066ms | http |
| 120.232.242.119:22222 | ✓ 1039ms | ✓ 1309ms | ✓ 941ms | 否 | ✓ 899ms | http |
| 45.140.147.155:1081 | ✓ 720ms | ✓ 1931ms | ✓ 1443ms | ✓ 1437ms | ✓ 1381ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1201ms | ✓ 1029ms | ✓ 1387ms | 否 | http |
| 178.236.245.17:3128 | ✓ 1120ms | 否 | ✓ 720ms | ✓ 1950ms | ✓ 1595ms | http |
| 61.72.221.194:3128 | ✓ 1510ms | 否 | ✓ 1297ms | ✓ 1041ms | ✓ 845ms | http |
| 61.72.110.94:3128 | ✓ 1514ms | 否 | ✓ 1545ms | 否 | ✓ 1652ms | http |
| 138.124.53.25:7443 | ✓ 651ms | ✓ 1740ms | 否 | 否 | ✓ 1491ms | http |
| 120.238.159.229:22222 | ✓ 979ms | ✓ 1356ms | ✓ 891ms | ✓ 1309ms | ✓ 1062ms | http |
| 103.84.95.54:7890 | ✓ 678ms | 否 | ✓ 639ms | ✓ 1088ms | ✓ 663ms | http |
| 162.248.165.72:1080 | ✓ 952ms | 否 | ✓ 1878ms | 否 | ✓ 1640ms | http |
| 14.225.222.247:7890 | ✓ 1276ms | 否 | 否 | ✓ 1018ms | ✓ 821ms | http |
| 14.56.107.244:3128 | ✓ 1215ms | 否 | 否 | ✓ 934ms | ✓ 858ms | http |
| 61.76.95.217:40088 | ✓ 947ms | ✓ 1881ms | ✓ 1145ms | ✓ 1199ms | ✓ 1057ms | http |
| 121.128.121.54:3128 | ✓ 610ms | ✓ 939ms | ✓ 945ms | ✓ 1620ms | ✓ 1837ms | http |
| 160.179.227.52:17464 | ✓ 1345ms | 否 | 否 | ✓ 1964ms | ✓ 1931ms | http |
| 150.107.140.238:3128 | ✓ 1668ms | 否 | 否 | ✓ 1189ms | ✓ 982ms | http |
| 101.43.255.96:80 | ✓ 1042ms | 否 | ✓ 1272ms | ✓ 1620ms | ✓ 1698ms | http |
| 81.70.169.194:80 | ✓ 1112ms | 否 | ✓ 1363ms | ✓ 1313ms | 否 | http |
| 61.72.221.234:3128 | ✓ 746ms | ✓ 1767ms | 否 | ✓ 1100ms | 否 | http |
| 136.49.39.94:8888 | 否 | ✓ 1402ms | ✓ 1136ms | 否 | ✓ 1105ms | http |
| 159.89.31.62:8080 | ✓ 1144ms | 否 | 否 | ✓ 1994ms | ✓ 1646ms | http |
| 91.233.223.147:3128 | ✓ 1195ms | 否 | ✓ 1636ms | 否 | ✓ 1707ms | http |
| 178.236.245.59:3128 | ✓ 1187ms | ✓ 1941ms | ✓ 1617ms | 否 | ✓ 1650ms | http |
| 190.9.109.205:999 | ✓ 923ms | 否 | ✓ 1267ms | ✓ 1628ms | ✓ 1171ms | http |
| 190.9.109.199:999 | ✓ 1613ms | 否 | ✓ 1157ms | ✓ 1334ms | 否 | http |
| 113.176.92.71:3128 | ✓ 897ms | ✓ 1599ms | ✓ 1179ms | ✓ 1185ms | ✓ 980ms | http |
| 46.249.103.192:443 | ✓ 726ms | 否 | ✓ 1448ms | ✓ 1548ms | 否 | http |
| 120.240.35.161:22222 | ✓ 1355ms | ✓ 1170ms | ✓ 992ms | ✓ 1194ms | ✓ 933ms | http |
| 120.240.35.160:22222 | ✓ 1212ms | ✓ 1292ms | ✓ 1228ms | ✓ 1088ms | ✓ 921ms | http |
| 120.240.35.178:22222 | ✓ 1987ms | ✓ 1236ms | ✓ 993ms | ✓ 1212ms | ✓ 890ms | http |
| 120.198.141.79:22222 | ✓ 933ms | ✓ 1229ms | ✓ 1251ms | ✓ 1288ms | 否 | http |
| 113.59.32.161:22222 | 否 | ✓ 1503ms | ✓ 1171ms | ✓ 1768ms | ✓ 1231ms | http |
| 167.172.69.123:8080 | ✓ 737ms | 否 | ✓ 1124ms | ✓ 1219ms | ✓ 840ms | http |
| 167.172.69.123:80 | ✓ 735ms | 否 | ✓ 1808ms | ✓ 1194ms | ✓ 913ms | http |
| 175.194.173.105:3128 | ✓ 746ms | ✓ 1181ms | ✓ 1626ms | 否 | 否 | http |
| 222.184.48.242:22222 | ✓ 1056ms | 否 | ✓ 1097ms | 否 | ✓ 1660ms | http |
| 188.132.141.249:443 | ✓ 1335ms | 否 | ✓ 1737ms | 否 | ✓ 1880ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1457ms | ✓ 1953ms | 否 | ✓ 979ms | http |
| 193.168.173.136:443 | ✓ 1500ms | 否 | ✓ 1215ms | 否 | ✓ 1592ms | http |
| 35.225.22.61:80 | 否 | ✓ 1288ms | ✓ 509ms | ✓ 1172ms | ✓ 971ms | http |
| 61.72.221.94:3128 | ✓ 848ms | ✓ 1542ms | ✓ 1901ms | 否 | ✓ 814ms | http |
| 91.107.175.112:10801 | ✓ 578ms | ✓ 1941ms | 否 | 否 | ✓ 1546ms | http |
| 222.184.48.236:22222 | ✓ 1018ms | ✓ 1425ms | ✓ 846ms | ✓ 1284ms | ✓ 949ms | http |
| 85.9.195.140:1080 | 否 | 否 | ✓ 697ms | ✓ 1825ms | ✓ 1837ms | http |
| 37.187.109.70:10111 | ✓ 1908ms | 否 | ✓ 994ms | 否 | ✓ 1733ms | http |
| 211.171.114.154:3128 | 否 | 否 | ✓ 1310ms | ✓ 1113ms | ✓ 1091ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1666ms | 否 | ✓ 1475ms | ✓ 859ms | http |
| 113.59.32.141:22222 | ✓ 1151ms | ✓ 1437ms | 否 | 否 | ✓ 1063ms | http |
| 168.235.110.63:3128 | ✓ 1652ms | 否 | ✓ 345ms | 否 | ✓ 1044ms | http |
| 34.101.184.164:3128 | ✓ 851ms | 否 | ✓ 1439ms | 否 | ✓ 1099ms | http |
| 91.193.240.157:9877 | ✓ 1202ms | 否 | ✓ 1106ms | 否 | ✓ 1989ms | http |
| 4.213.222.169:3128 | ✓ 1974ms | 否 | ✓ 1540ms | ✓ 1858ms | ✓ 1602ms | http |
| 14.225.217.30:7890 | 否 | ✓ 1893ms | ✓ 835ms | ✓ 1142ms | ✓ 1348ms | http |
| 117.159.239.50:22222 | 否 | ✓ 1009ms | ✓ 798ms | ✓ 1082ms | ✓ 874ms | http |
| 117.159.239.52:22222 | ✓ 871ms | ✓ 1054ms | ✓ 752ms | ✓ 1058ms | ✓ 840ms | http |
| 120.198.141.75:22222 | ✓ 1912ms | ✓ 1306ms | ✓ 926ms | ✓ 1166ms | ✓ 1011ms | http |
| 120.240.35.176:22222 | ✓ 970ms | ✓ 1231ms | ✓ 1072ms | ✓ 1104ms | ✓ 854ms | http |
| 103.139.138.194:3128 | ✓ 1843ms | 否 | 否 | ✓ 1485ms | ✓ 1153ms | http |
| 121.230.9.161:1080 | 否 | ✓ 1477ms | 否 | ✓ 1695ms | ✓ 1249ms | http |
| 222.184.48.252:22222 | ✓ 1684ms | ✓ 1872ms | ✓ 973ms | ✓ 1915ms | ✓ 1597ms | http |
| 222.184.48.235:22222 | 否 | ✓ 1222ms | ✓ 1746ms | 否 | ✓ 1033ms | http |
| 120.198.141.84:22222 | 否 | ✓ 1626ms | ✓ 1283ms | ✓ 1595ms | 否 | http |
| 14.225.222.185:7890 | 否 | ✓ 1449ms | 否 | ✓ 1200ms | ✓ 899ms | http |
| 210.223.44.230:3128 | ✓ 1290ms | 否 | 否 | ✓ 1192ms | ✓ 1601ms | http |
| 45.136.198.40:3128 | ✓ 1449ms | 否 | ✓ 1553ms | 否 | ✓ 1854ms | http |
| 172.212.68.37:3128 | ✓ 714ms | 否 | 否 | ✓ 1806ms | ✓ 1259ms | http |
| 103.82.23.118:5247 | 否 | 否 | ✓ 1824ms | ✓ 1712ms | ✓ 1551ms | http |
| 113.59.32.162:22222 | ✓ 1280ms | ✓ 1612ms | ✓ 1388ms | ✓ 1783ms | ✓ 1301ms | http |
| 14.225.222.164:7890 | ✓ 1580ms | 否 | ✓ 1377ms | 否 | ✓ 1912ms | http |
| 217.77.102.18:3128 | ✓ 1658ms | 否 | ✓ 1584ms | 否 | ✓ 1818ms | http |
| 46.183.25.8:443 | ✓ 970ms | 否 | ✓ 1849ms | ✓ 781ms | 否 | http |
| 105.158.244.76:18826 | ✓ 1320ms | 否 | ✓ 1701ms | 否 | ✓ 1729ms | http |
| 66.54.106.56:8080 | ✓ 1014ms | 否 | ✓ 1862ms | ✓ 1876ms | 否 | http |
| 1.225.116.115:1080 | 否 | 否 | ✓ 1008ms | ✓ 1415ms | ✓ 889ms | http |
| 69.48.179.20:3128 | ✓ 397ms | 否 | ✓ 290ms | ✓ 1490ms | 否 | http |
| 103.215.36.88:18977 | ✓ 1913ms | 否 | ✓ 1020ms | ✓ 1568ms | ✓ 1004ms | http |
| 183.249.5.214:22222 | ✓ 746ms | ✓ 993ms | ✓ 841ms | 否 | ✓ 1900ms | http |
| 120.240.35.177:22222 | ✓ 1291ms | ✓ 1336ms | ✓ 1097ms | ✓ 1332ms | ✓ 1113ms | http |
| 62.113.119.14:8080 | ✓ 1835ms | 否 | ✓ 1189ms | ✓ 1711ms | ✓ 1253ms | http |
| 113.59.32.142:22222 | ✓ 1594ms | ✓ 1939ms | ✓ 1381ms | ✓ 1865ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1242ms | ✓ 1975ms | ✓ 1833ms | ✓ 1657ms | 否 | http |
| 125.128.12.144:3128 | 否 | 否 | ✓ 963ms | ✓ 1412ms | ✓ 824ms | http |
| 125.128.12.14:3128 | ✓ 1728ms | ✓ 1353ms | ✓ 636ms | ✓ 1116ms | ✓ 792ms | http |
| 154.64.240.39:1080 | ✓ 1577ms | ✓ 1776ms | 否 | 否 | ✓ 1894ms | http |
| 159.223.42.219:3128 | 否 | 否 | ✓ 1442ms | ✓ 1273ms | ✓ 1554ms | http |
| 157.20.253.164:8080 | 否 | 否 | ✓ 1187ms | ✓ 1689ms | ✓ 1670ms | http |

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
