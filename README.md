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

最后更新：2026-04-23 22:42:24 UTC（2026-04-24 06:42:24 UTC+8）

**代理总数：57**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 57 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 57 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 47.85.51.197:1080 | ✓ 589ms | 否 | ✓ 289ms | ✓ 882ms | ✓ 659ms | http |
| 46.101.95.183:8888 | ✓ 631ms | ✓ 1713ms | ✓ 625ms | ✓ 1435ms | ✓ 1172ms | http |
| 1.231.81.166:3128 | ✓ 1009ms | ✓ 1232ms | ✓ 1185ms | ✓ 1105ms | ✓ 894ms | http |
| 113.160.132.26:8080 | ✓ 1701ms | ✓ 1522ms | 否 | ✓ 1493ms | ✓ 1308ms | http |
| 212.58.132.5:8888 | ✓ 1862ms | 否 | ✓ 1825ms | ✓ 1552ms | ✓ 1262ms | http |
| 115.231.181.40:8128 | ✓ 1690ms | ✓ 1572ms | ✓ 1114ms | 否 | ✓ 1168ms | http |
| 35.225.22.61:80 | ✓ 1032ms | ✓ 1289ms | ✓ 1063ms | ✓ 1126ms | 否 | http |
| 62.113.119.14:8080 | ✓ 591ms | ✓ 1595ms | ✓ 1136ms | ✓ 1422ms | ✓ 1090ms | http |
| 218.108.131.186:17890 | ✓ 1087ms | 否 | ✓ 981ms | ✓ 1366ms | 否 | http |
| 34.96.238.40:8080 | ✓ 1434ms | ✓ 1316ms | ✓ 1210ms | ✓ 1377ms | ✓ 1228ms | http |
| 84.47.150.125:1080 | ✓ 1235ms | 否 | ✓ 1954ms | 否 | ✓ 1753ms | http |
| 130.61.174.200:1080 | ✓ 874ms | ✓ 1586ms | 否 | ✓ 1288ms | 否 | http |
| 120.92.108.86:7890 | ✓ 1685ms | 否 | 否 | ✓ 1912ms | ✓ 1553ms | http |
| 59.46.216.131:30001 | ✓ 1169ms | ✓ 1559ms | 否 | ✓ 1597ms | 否 | http |
| 177.93.132.244:3128 | ✓ 1030ms | 否 | ✓ 927ms | 否 | ✓ 1736ms | http |
| 91.99.15.45:2095 | ✓ 640ms | ✓ 1502ms | ✓ 1631ms | 否 | ✓ 1796ms | http |
| 45.153.231.229:8080 | ✓ 1033ms | 否 | ✓ 1828ms | 否 | ✓ 1986ms | http |
| 34.71.229.255:3128 | ✓ 484ms | ✓ 1437ms | ✓ 989ms | ✓ 1139ms | ✓ 1140ms | http |
| 168.110.52.228:3128 | ✓ 771ms | 否 | ✓ 865ms | ✓ 1031ms | ✓ 948ms | http |
| 45.129.141.143:3128 | ✓ 673ms | ✓ 1947ms | ✓ 1504ms | 否 | ✓ 1384ms | http |
| 38.180.2.107:3128 | ✓ 1054ms | ✓ 1946ms | ✓ 1627ms | 否 | ✓ 1990ms | http |
| 120.92.212.16:8890 | ✓ 1075ms | ✓ 1457ms | ✓ 1091ms | 否 | ✓ 1077ms | http |
| 137.59.47.73:3128 | ✓ 1918ms | 否 | ✓ 1514ms | 否 | ✓ 1187ms | http |
| 85.190.99.143:443 | ✓ 518ms | ✓ 1856ms | ✓ 1215ms | ✓ 1895ms | 否 | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1001ms | ✓ 1417ms | ✓ 1108ms | http |
| 120.92.212.16:7890 | ✓ 1110ms | 否 | ✓ 1135ms | ✓ 1400ms | ✓ 1109ms | http |
| 162.240.154.26:3128 | ✓ 656ms | ✓ 1055ms | ✓ 1008ms | ✓ 1058ms | ✓ 989ms | http |
| 8.211.166.184:8081 | ✓ 1728ms | ✓ 1101ms | ✓ 934ms | ✓ 1036ms | 否 | http |
| 223.84.151.86:30005 | ✓ 1811ms | ✓ 1536ms | 否 | 否 | ✓ 1777ms | http |
| 45.140.147.82:1081 | ✓ 699ms | ✓ 1455ms | ✓ 1602ms | ✓ 1703ms | ✓ 1216ms | http |
| 45.76.207.177:40000 | ✓ 827ms | 否 | ✓ 1940ms | 否 | ✓ 898ms | http |
| 149.56.24.51:3128 | 否 | ✓ 1261ms | 否 | ✓ 1469ms | ✓ 1340ms | http |
| 114.237.77.202:1080 | ✓ 1122ms | ✓ 1400ms | ✓ 1053ms | ✓ 1374ms | ✓ 1213ms | http |
| 103.187.146.151:3128 | ✓ 1582ms | 否 | ✓ 1274ms | ✓ 1478ms | ✓ 1061ms | http |
| 121.230.9.11:1080 | ✓ 1274ms | ✓ 1535ms | 否 | ✓ 1723ms | ✓ 1182ms | http |
| 168.222.254.136:8888 | ✓ 1937ms | 否 | ✓ 1974ms | 否 | ✓ 1856ms | http |
| 91.193.240.157:9877 | ✓ 1090ms | 否 | ✓ 914ms | 否 | ✓ 1863ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 835ms | ✓ 1289ms | ✓ 1012ms | http |
| 152.32.132.190:7890 | 否 | ✓ 1895ms | ✓ 1339ms | ✓ 1866ms | 否 | http |
| 121.230.8.136:1080 | ✓ 1387ms | ✓ 1745ms | ✓ 1324ms | ✓ 1893ms | ✓ 1386ms | http |
| 77.110.113.24:40000 | ✓ 903ms | 否 | ✓ 569ms | ✓ 1943ms | ✓ 1450ms | http |
| 20.164.75.153:8080 | ✓ 1803ms | 否 | ✓ 1152ms | 否 | ✓ 1881ms | http |
| 47.84.131.156:8100 | ✓ 1957ms | ✓ 1868ms | ✓ 1017ms | ✓ 1262ms | ✓ 1118ms | http |
| 92.113.149.172:8080 | ✓ 1117ms | ✓ 1689ms | ✓ 1212ms | 否 | 否 | http |
| 45.186.6.104:3128 | ✓ 1912ms | ✓ 1918ms | ✓ 1683ms | 否 | 否 | http |
| 210.45.76.58:42992 | 否 | 否 | ✓ 1333ms | ✓ 1513ms | ✓ 1244ms | http |
| 168.144.75.9:3128 | ✓ 1311ms | 否 | ✓ 1948ms | 否 | ✓ 1634ms | http |
| 13.51.196.44:443 | 否 | 否 | ✓ 1143ms | ✓ 1829ms | ✓ 1582ms | http |
| 124.121.2.131:8080 | 否 | 否 | ✓ 1599ms | ✓ 1822ms | ✓ 1863ms | http |
| 208.87.243.199:7878 | 否 | ✓ 1330ms | ✓ 995ms | ✓ 1364ms | 否 | http |
| 220.197.44.36:3128 | 否 | 否 | ✓ 1192ms | ✓ 1727ms | ✓ 1283ms | http |
| 45.129.141.174:3128 | ✓ 1124ms | 否 | 否 | ✓ 1880ms | ✓ 1709ms | http |
| 8.219.195.129:1080 | ✓ 1895ms | ✓ 1943ms | ✓ 1038ms | ✓ 1254ms | ✓ 1057ms | http |
| 118.113.246.172:1080 | ✓ 1361ms | ✓ 1665ms | ✓ 1323ms | ✓ 1864ms | ✓ 1366ms | http |
| 152.42.177.32:8888 | ✓ 1180ms | 否 | ✓ 1162ms | ✓ 1530ms | ✓ 1498ms | http |
| 104.248.151.93:9090 | ✓ 1955ms | 否 | ✓ 1082ms | ✓ 1429ms | ✓ 1240ms | http |
| 23.95.76.201:8443 | ✓ 553ms | ✓ 1359ms | ✓ 844ms | 否 | 否 | http |

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
