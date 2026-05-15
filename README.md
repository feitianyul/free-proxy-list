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

最后更新：2026-05-15 16:14:19 UTC（2026-05-16 00:14:19 UTC+8）

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
| 1.231.81.166:3128 | 否 | ✓ 1094ms | ✓ 1341ms | ✓ 1259ms | ✓ 959ms | http |
| 137.59.47.73:3128 | 否 | ✓ 1479ms | ✓ 1459ms | ✓ 1236ms | ✓ 1459ms | http |
| 218.108.131.186:17890 | ✓ 1090ms | ✓ 1195ms | ✓ 878ms | ✓ 1215ms | ✓ 975ms | http |
| 115.231.181.40:8128 | ✓ 1001ms | ✓ 1238ms | ✓ 953ms | ✓ 1319ms | ✓ 1107ms | http |
| 91.217.81.131:1080 | ✓ 767ms | 否 | ✓ 1179ms | 否 | ✓ 1624ms | http |
| 113.160.132.26:8080 | ✓ 1969ms | 否 | ✓ 1388ms | 否 | ✓ 1149ms | http |
| 91.242.229.129:8092 | ✓ 1468ms | 否 | ✓ 1665ms | 否 | ✓ 1767ms | http |
| 185.191.236.162:3128 | 否 | 否 | ✓ 929ms | ✓ 1701ms | ✓ 1176ms | http |
| 59.46.216.131:30001 | 否 | 否 | ✓ 1245ms | ✓ 1822ms | ✓ 1219ms | http |
| 86.104.72.219:1082 | ✓ 814ms | ✓ 1081ms | ✓ 1066ms | ✓ 1052ms | ✓ 945ms | http |
| 86.104.72.219:1081 | ✓ 812ms | 否 | ✓ 147ms | ✓ 1054ms | ✓ 965ms | http |
| 128.199.116.219:9090 | 否 | 否 | ✓ 1256ms | ✓ 1181ms | ✓ 911ms | http |
| 128.199.113.85:9090 | ✓ 973ms | 否 | ✓ 1003ms | ✓ 1156ms | ✓ 942ms | http |
| 128.199.254.13:9090 | ✓ 974ms | 否 | ✓ 1000ms | ✓ 1158ms | ✓ 943ms | http |
| 103.157.200.126:3128 | ✓ 1271ms | 否 | ✓ 1252ms | ✓ 1888ms | ✓ 1489ms | http |
| 5.252.33.13:2025 | ✓ 1419ms | 否 | ✓ 1489ms | 否 | ✓ 1834ms | http |
| 34.71.229.255:3128 | 否 | 否 | ✓ 1140ms | ✓ 1743ms | ✓ 1493ms | http |
| 120.92.212.16:7890 | ✓ 1471ms | ✓ 1931ms | 否 | ✓ 1900ms | 否 | http |
| 42.200.76.16:3888 | ✓ 997ms | 否 | ✓ 1244ms | ✓ 1035ms | ✓ 1056ms | http |
| 45.88.0.113:3128 | 否 | 否 | ✓ 516ms | ✓ 1422ms | ✓ 1127ms | http |
| 160.238.65.2:3128 | 否 | ✓ 1971ms | ✓ 1559ms | 否 | ✓ 1656ms | http |
| 84.47.150.125:1080 | ✓ 803ms | 否 | ✓ 1744ms | 否 | ✓ 1690ms | http |
| 128.199.114.189:9090 | ✓ 1305ms | 否 | ✓ 1233ms | ✓ 1251ms | ✓ 973ms | http |
| 119.28.51.157:3128 | ✓ 965ms | ✓ 1215ms | ✓ 990ms | ✓ 1094ms | ✓ 1087ms | http |
| 120.92.212.16:8890 | ✓ 1106ms | 否 | ✓ 1747ms | ✓ 1609ms | ✓ 1026ms | http |
| 157.0.142.246:10057 | 否 | 否 | ✓ 1242ms | ✓ 1452ms | ✓ 1625ms | http |
| 166.88.55.83:7890 | ✓ 722ms | ✓ 1223ms | ✓ 733ms | ✓ 924ms | ✓ 739ms | http |
| 103.21.220.141:3128 | ✓ 830ms | 否 | ✓ 912ms | ✓ 1008ms | ✓ 861ms | http |
| 129.80.217.21:444 | 否 | 否 | ✓ 1360ms | ✓ 1115ms | ✓ 1321ms | http |
| 34.96.238.40:8080 | ✓ 1903ms | ✓ 1693ms | ✓ 1329ms | ✓ 1095ms | 否 | http |
| 8.154.21.175:3128 | ✓ 1038ms | ✓ 1121ms | ✓ 1983ms | ✓ 1442ms | ✓ 992ms | http |
| 212.58.132.5:8888 | ✓ 1780ms | 否 | ✓ 1116ms | ✓ 1407ms | ✓ 1147ms | http |
| 121.230.8.97:1080 | 否 | ✓ 1572ms | ✓ 1548ms | 否 | ✓ 1085ms | http |
| 113.176.92.71:3128 | 否 | ✓ 1533ms | 否 | ✓ 1338ms | ✓ 1523ms | http |
| 222.127.77.167:8085 | 否 | 否 | ✓ 1866ms | ✓ 1651ms | ✓ 1722ms | http |
| 104.248.151.93:9090 | ✓ 813ms | 否 | ✓ 842ms | ✓ 1155ms | ✓ 919ms | http |
| 121.130.177.28:8888 | ✓ 1141ms | 否 | 否 | ✓ 1688ms | ✓ 1296ms | http |
| 103.134.85.145:3128 | ✓ 925ms | 否 | ✓ 922ms | ✓ 1383ms | ✓ 1039ms | http |
| 45.144.28.116:3128 | ✓ 849ms | 否 | ✓ 547ms | ✓ 1726ms | 否 | http |
| 160.238.65.3:3128 | ✓ 1967ms | 否 | ✓ 677ms | ✓ 1402ms | ✓ 1425ms | http |
| 47.117.135.38:3128 | ✓ 1768ms | ✓ 1957ms | ✓ 1668ms | ✓ 1633ms | ✓ 1244ms | http |
| 3.101.133.120:80 | ✓ 432ms | ✓ 1630ms | ✓ 1308ms | ✓ 1315ms | ✓ 949ms | http |
| 185.21.15.206:3128 | 否 | 否 | ✓ 1632ms | ✓ 1565ms | ✓ 1577ms | http |
| 110.172.29.131:3128 | ✓ 1548ms | 否 | ✓ 1355ms | ✓ 1384ms | ✓ 1603ms | http |
| 210.223.44.230:3128 | ✓ 1573ms | ✓ 1003ms | ✓ 809ms | ✓ 989ms | ✓ 777ms | http |
| 94.241.169.176:1080 | 否 | ✓ 1956ms | ✓ 1551ms | 否 | ✓ 1946ms | http |
| 146.190.80.158:9090 | 否 | 否 | ✓ 1671ms | ✓ 1302ms | ✓ 939ms | http |
| 128.199.121.61:9090 | ✓ 811ms | 否 | ✓ 1788ms | ✓ 1252ms | ✓ 1224ms | http |
| 152.42.170.187:9090 | ✓ 831ms | 否 | ✓ 1353ms | ✓ 1211ms | ✓ 1052ms | http |
| 168.222.254.136:8888 | ✓ 1477ms | ✓ 1861ms | ✓ 1881ms | ✓ 1882ms | 否 | http |
| 152.32.132.190:7890 | 否 | ✓ 1993ms | 否 | ✓ 1870ms | ✓ 1618ms | http |
| 152.70.91.193:40000 | ✓ 1214ms | 否 | 否 | ✓ 1260ms | ✓ 1637ms | http |
| 8.219.97.248:80 | ✓ 1442ms | 否 | ✓ 1224ms | ✓ 1912ms | 否 | http |
| 209.126.84.232:8888 | ✓ 1901ms | 否 | ✓ 1821ms | 否 | ✓ 1694ms | http |
| 57.129.144.178:40000 | ✓ 820ms | 否 | ✓ 1990ms | ✓ 1812ms | ✓ 1528ms | http |
| 160.238.65.4:3128 | ✓ 1855ms | 否 | ✓ 1653ms | 否 | ✓ 1146ms | http |
| 107.175.85.198:1080 | ✓ 1203ms | 否 | ✓ 837ms | ✓ 1448ms | ✓ 1016ms | http |
| 61.52.131.172:8443 | ✓ 965ms | ✓ 1399ms | ✓ 1119ms | ✓ 1282ms | ✓ 1069ms | http |
| 103.172.70.173:8080 | ✓ 1468ms | 否 | ✓ 1773ms | ✓ 1572ms | ✓ 1448ms | http |
| 103.157.117.116:8080 | 否 | 否 | ✓ 1586ms | ✓ 1819ms | ✓ 1871ms | http |

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
