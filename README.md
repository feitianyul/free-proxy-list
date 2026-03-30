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

最后更新：2026-03-30 14:38:45 UTC（2026-03-30 22:38:45 UTC+8）

**代理总数：60**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 60 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 60 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 39.185.46.193:5911 | ✓ 737ms | ✓ 777ms | ✓ 629ms | ✓ 976ms | ✓ 648ms | http |
| 43.99.54.236:5555 | ✓ 770ms | ✓ 1134ms | ✓ 621ms | ✓ 800ms | ✓ 631ms | http |
| 147.161.210.140:8800 | ✓ 1367ms | 否 | ✓ 949ms | ✓ 964ms | ✓ 1037ms | http |
| 167.103.115.102:8800 | ✓ 1570ms | 否 | ✓ 929ms | 否 | ✓ 1142ms | http |
| 147.161.239.240:8800 | ✓ 1445ms | ✓ 1979ms | ✓ 1384ms | ✓ 1830ms | ✓ 1570ms | http |
| 113.160.132.26:8080 | ✓ 1674ms | 否 | ✓ 1377ms | ✓ 1608ms | 否 | http |
| 167.103.34.108:8800 | ✓ 1663ms | 否 | ✓ 1462ms | ✓ 1574ms | ✓ 1564ms | http |
| 95.213.217.168:52004 | ✓ 1465ms | 否 | ✓ 1146ms | 否 | ✓ 1956ms | http |
| 101.47.73.135:3128 | ✓ 1580ms | 否 | 否 | ✓ 1559ms | ✓ 1945ms | http |
| 183.249.5.117:22222 | ✓ 707ms | 否 | ✓ 863ms | ✓ 1163ms | ✓ 811ms | http |
| 183.249.5.105:22222 | ✓ 830ms | ✓ 870ms | ✓ 1122ms | 否 | ✓ 746ms | http |
| 167.103.144.127:8800 | ✓ 1605ms | 否 | ✓ 1277ms | 否 | ✓ 1432ms | http |
| 1.231.81.166:3128 | ✓ 1126ms | ✓ 1257ms | ✓ 1831ms | ✓ 1536ms | ✓ 1270ms | http |
| 167.103.31.122:8800 | ✓ 1644ms | 否 | ✓ 1274ms | ✓ 1662ms | ✓ 1544ms | http |
| 150.107.140.238:3128 | ✓ 1589ms | 否 | ✓ 1938ms | 否 | ✓ 973ms | http |
| 222.184.48.251:22222 | ✓ 905ms | ✓ 1886ms | ✓ 1553ms | ✓ 1288ms | 否 | http |
| 59.46.216.131:30001 | ✓ 927ms | ✓ 1313ms | ✓ 1070ms | 否 | 否 | http |
| 110.52.243.227:10809 | ✓ 922ms | ✓ 1203ms | ✓ 1037ms | 否 | ✓ 998ms | http |
| 120.92.212.16:8890 | ✓ 928ms | ✓ 1181ms | 否 | ✓ 1156ms | ✓ 920ms | http |
| 42.96.16.158:1311 | ✓ 1489ms | 否 | ✓ 1932ms | 否 | ✓ 1398ms | http |
| 101.43.127.100:8877 | 否 | ✓ 1675ms | ✓ 1818ms | ✓ 1969ms | ✓ 1578ms | http |
| 210.223.44.230:3128 | ✓ 616ms | 否 | 否 | ✓ 873ms | ✓ 983ms | http |
| 35.225.22.61:80 | ✓ 435ms | 否 | ✓ 1098ms | ✓ 1179ms | ✓ 1189ms | http |
| 115.231.181.40:8128 | 否 | 否 | ✓ 910ms | ✓ 1164ms | ✓ 1573ms | http |
| 5.104.87.17:8051 | ✓ 1434ms | 否 | ✓ 1281ms | ✓ 1423ms | ✓ 1222ms | http |
| 209.126.84.232:8888 | ✓ 1691ms | ✓ 1687ms | 否 | ✓ 1296ms | 否 | http |
| 20.78.118.91:8561 | 否 | 否 | ✓ 1630ms | ✓ 1725ms | ✓ 1180ms | http |
| 177.234.217.88:999 | ✓ 1532ms | 否 | ✓ 1047ms | ✓ 1873ms | 否 | http |
| 20.78.26.206:8561 | ✓ 503ms | ✓ 785ms | ✓ 438ms | ✓ 861ms | ✓ 681ms | http |
| 103.84.95.54:7890 | ✓ 959ms | 否 | ✓ 1305ms | 否 | ✓ 748ms | http |
| 45.140.147.155:1081 | ✓ 1662ms | 否 | ✓ 888ms | ✓ 1414ms | 否 | http |
| 120.92.212.16:7890 | ✓ 974ms | ✓ 1153ms | ✓ 876ms | ✓ 1164ms | ✓ 914ms | http |
| 101.32.244.83:8080 | ✓ 1003ms | ✓ 1406ms | ✓ 894ms | ✓ 1126ms | ✓ 1159ms | http |
| 121.43.196.213:8222 | ✓ 879ms | ✓ 982ms | ✓ 894ms | ✓ 1044ms | ✓ 866ms | http |
| 121.43.196.210:8222 | ✓ 918ms | ✓ 983ms | ✓ 839ms | ✓ 1050ms | ✓ 912ms | http |
| 114.55.226.123:10086 | ✓ 1042ms | ✓ 1849ms | ✓ 970ms | ✓ 1277ms | ✓ 1002ms | http |
| 183.249.5.111:22222 | ✓ 1281ms | ✓ 1380ms | ✓ 789ms | ✓ 952ms | 否 | http |
| 183.249.5.110:22222 | 否 | ✓ 990ms | ✓ 719ms | ✓ 962ms | ✓ 804ms | http |
| 38.34.179.100:8452 | ✓ 1848ms | 否 | 否 | ✓ 754ms | ✓ 760ms | http |
| 195.123.209.48:3128 | ✓ 1159ms | ✓ 1927ms | ✓ 1619ms | 否 | 否 | http |
| 121.230.8.237:1080 | ✓ 1461ms | ✓ 1523ms | ✓ 1889ms | ✓ 1885ms | 否 | http |
| 91.233.223.147:3128 | ✓ 1064ms | 否 | ✓ 1905ms | 否 | ✓ 1786ms | http |
| 222.184.48.252:22222 | ✓ 939ms | ✓ 1105ms | 否 | 否 | ✓ 872ms | http |
| 222.184.48.242:22222 | ✓ 940ms | 否 | ✓ 938ms | ✓ 1384ms | ✓ 882ms | http |
| 103.113.70.189:1081 | ✓ 1322ms | 否 | ✓ 551ms | 否 | ✓ 1909ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1592ms | ✓ 1577ms | ✓ 1222ms | http |
| 45.136.198.40:3128 | ✓ 905ms | ✓ 1891ms | ✓ 1289ms | ✓ 1952ms | ✓ 1734ms | http |
| 45.12.151.226:2829 | ✓ 1232ms | 否 | 否 | ✓ 1843ms | ✓ 1466ms | http |
| 45.81.131.195:8888 | ✓ 324ms | ✓ 1616ms | ✓ 77ms | ✓ 1260ms | ✓ 1703ms | http |
| 62.171.161.88:2018 | ✓ 1278ms | ✓ 1900ms | ✓ 1026ms | ✓ 1796ms | ✓ 1663ms | http |
| 142.171.157.207:3128 | ✓ 253ms | ✓ 1888ms | ✓ 826ms | ✓ 901ms | ✓ 661ms | http |
| 222.184.48.241:22222 | ✓ 1120ms | 否 | ✓ 1968ms | 否 | ✓ 1262ms | http |
| 118.31.1.154:80 | 否 | 否 | ✓ 1527ms | ✓ 1165ms | ✓ 847ms | http |
| 168.110.52.228:3128 | ✓ 1919ms | 否 | ✓ 1249ms | 否 | ✓ 1821ms | http |
| 160.238.65.9:3128 | ✓ 1969ms | ✓ 1696ms | ✓ 1622ms | 否 | 否 | http |
| 45.144.28.81:10808 | ✓ 767ms | 否 | 否 | ✓ 1679ms | ✓ 1698ms | http |
| 220.197.44.36:3128 | ✓ 1666ms | 否 | ✓ 1483ms | 否 | ✓ 1623ms | http |
| 103.39.51.190:8080 | ✓ 1289ms | 否 | 否 | ✓ 1764ms | ✓ 1471ms | http |
| 217.76.245.80:999 | 否 | 否 | ✓ 1534ms | ✓ 1419ms | ✓ 1203ms | http |
| 103.52.114.95:3128 | ✓ 1627ms | 否 | ✓ 783ms | ✓ 1221ms | ✓ 981ms | http |

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
