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

最后更新：2026-05-02 10:53:12 UTC（2026-05-02 18:53:12 UTC+8）

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
| 206.206.126.177:2412 | ✓ 1040ms | 否 | ✓ 1018ms | ✓ 1214ms | 否 | http |
| 1.231.81.166:3128 | 否 | ✓ 1457ms | ✓ 819ms | ✓ 1494ms | 否 | http |
| 46.101.95.183:8888 | ✓ 1399ms | 否 | ✓ 1584ms | ✓ 1885ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1169ms | 否 | ✓ 1821ms | ✓ 1523ms | ✓ 1245ms | http |
| 113.160.132.26:8080 | 否 | 否 | ✓ 1193ms | ✓ 1560ms | ✓ 1292ms | http |
| 45.167.124.71:999 | ✓ 1067ms | ✓ 1777ms | ✓ 1609ms | ✓ 1681ms | ✓ 1420ms | http |
| 86.104.72.219:1081 | ✓ 682ms | ✓ 928ms | ✓ 188ms | 否 | 否 | http |
| 103.157.200.126:3128 | ✓ 1943ms | 否 | ✓ 1123ms | ✓ 1965ms | ✓ 1353ms | http |
| 217.76.245.80:999 | ✓ 930ms | ✓ 1364ms | ✓ 1222ms | ✓ 1907ms | ✓ 1094ms | http |
| 45.153.231.229:8080 | ✓ 1334ms | 否 | ✓ 1000ms | 否 | ✓ 1552ms | http |
| 62.113.119.14:8080 | ✓ 509ms | ✓ 1599ms | ✓ 1937ms | ✓ 1469ms | ✓ 1136ms | http |
| 223.84.151.86:30005 | ✓ 1521ms | ✓ 1450ms | ✓ 1472ms | ✓ 1790ms | ✓ 1558ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1576ms | ✓ 1312ms | ✓ 1632ms | ✓ 1294ms | http |
| 168.222.254.136:8888 | ✓ 982ms | ✓ 1705ms | ✓ 1664ms | 否 | 否 | http |
| 91.184.241.12:443 | ✓ 1503ms | ✓ 1868ms | ✓ 1266ms | 否 | 否 | http |
| 107.150.41.226:18080 | ✓ 417ms | ✓ 995ms | ✓ 266ms | 否 | 否 | http |
| 47.85.51.197:1080 | ✓ 117ms | 否 | ✓ 500ms | ✓ 914ms | ✓ 828ms | http |
| 80.92.204.47:1081 | ✓ 1162ms | 否 | ✓ 667ms | ✓ 1601ms | ✓ 1887ms | http |
| 148.230.4.241:999 | ✓ 722ms | ✓ 1781ms | ✓ 646ms | 否 | 否 | http |
| 45.140.147.155:1082 | ✓ 966ms | 否 | ✓ 994ms | 否 | ✓ 1520ms | http |
| 149.51.42.10:8080 | ✓ 317ms | ✓ 975ms | 否 | ✓ 1989ms | 否 | http |
| 20.127.128.70:8080 | ✓ 1618ms | 否 | ✓ 1306ms | 否 | ✓ 1534ms | http |
| 45.129.141.143:3128 | ✓ 1460ms | 否 | ✓ 1696ms | 否 | ✓ 1524ms | http |
| 121.43.196.210:8222 | ✓ 1211ms | ✓ 1268ms | ✓ 1113ms | ✓ 1383ms | ✓ 1151ms | http |
| 72.11.150.178:6005 | ✓ 409ms | 否 | ✓ 852ms | 否 | ✓ 1009ms | http |
| 72.11.151.159:6005 | ✓ 908ms | ✓ 1283ms | ✓ 913ms | ✓ 1234ms | ✓ 1008ms | http |
| 34.101.184.164:3128 | ✓ 927ms | 否 | ✓ 974ms | ✓ 1392ms | ✓ 1155ms | http |
| 120.92.212.16:7890 | ✓ 1221ms | 否 | 否 | ✓ 1758ms | ✓ 1648ms | http |
| 181.78.44.63:999 | 否 | ✓ 1532ms | ✓ 1602ms | ✓ 1725ms | 否 | http |
| 86.104.74.110:1081 | ✓ 520ms | 否 | ✓ 660ms | ✓ 1515ms | ✓ 1039ms | http |
| 149.51.42.10:3128 | ✓ 221ms | ✓ 980ms | 否 | ✓ 1377ms | 否 | http |
| 121.135.144.141:8052 | ✓ 1765ms | ✓ 1314ms | ✓ 1050ms | ✓ 1249ms | ✓ 984ms | http |
| 183.238.3.150:7897 | ✓ 1308ms | ✓ 1363ms | ✓ 1128ms | ✓ 1351ms | ✓ 1079ms | http |
| 213.111.146.36:18080 | ✓ 1225ms | ✓ 1966ms | ✓ 1489ms | ✓ 1819ms | ✓ 1358ms | http |
| 185.195.71.218:18080 | ✓ 1586ms | 否 | ✓ 1376ms | 否 | ✓ 1447ms | http |
| 8.154.21.175:3128 | ✓ 996ms | 否 | 否 | ✓ 1293ms | ✓ 1140ms | http |
| 103.166.182.144:3128 | 否 | ✓ 1812ms | ✓ 1228ms | ✓ 1483ms | ✓ 1173ms | http |
| 107.174.80.186:3128 | ✓ 1614ms | 否 | 否 | ✓ 1208ms | ✓ 1031ms | http |
| 121.230.9.209:1080 | ✓ 1276ms | ✓ 1991ms | 否 | ✓ 1798ms | 否 | http |
| 38.130.38.97:8080 | ✓ 1431ms | ✓ 1986ms | 否 | ✓ 1543ms | ✓ 1684ms | http |
| 101.32.244.83:8080 | ✓ 1232ms | ✓ 1988ms | ✓ 1167ms | ✓ 1482ms | ✓ 1520ms | http |
| 121.43.196.213:8222 | ✓ 1104ms | ✓ 1306ms | ✓ 1022ms | ✓ 1346ms | ✓ 1151ms | http |
| 117.236.124.166:3128 | ✓ 1032ms | 否 | ✓ 1046ms | 否 | ✓ 1454ms | http |
| 38.188.247.12:999 | ✓ 1071ms | ✓ 1541ms | ✓ 1370ms | ✓ 1570ms | ✓ 1954ms | http |
| 20.164.75.153:8080 | ✓ 1150ms | 否 | ✓ 1884ms | 否 | ✓ 1785ms | http |
| 86.104.72.219:1082 | ✓ 1279ms | ✓ 876ms | ✓ 333ms | ✓ 1011ms | ✓ 700ms | http |
| 3.101.133.120:80 | ✓ 881ms | ✓ 1270ms | ✓ 1549ms | ✓ 1491ms | ✓ 776ms | http |
| 182.53.202.208:8080 | 否 | 否 | ✓ 1602ms | ✓ 1805ms | ✓ 1655ms | http |
| 34.96.238.40:8080 | ✓ 1662ms | ✓ 1991ms | 否 | ✓ 1789ms | 否 | http |
| 202.129.206.239:3128 | ✓ 1351ms | 否 | ✓ 1986ms | 否 | ✓ 1843ms | http |
| 167.160.191.204:6005 | ✓ 609ms | ✓ 1224ms | ✓ 814ms | ✓ 1744ms | ✓ 1966ms | http |
| 218.108.131.186:17890 | 否 | ✓ 1267ms | ✓ 992ms | ✓ 1282ms | ✓ 1058ms | http |
| 103.82.23.118:5230 | ✓ 1939ms | 否 | ✓ 1615ms | 否 | ✓ 1716ms | http |
| 152.32.132.190:7890 | ✓ 1664ms | 否 | ✓ 1204ms | ✓ 1363ms | ✓ 863ms | http |
| 103.3.58.162:8088 | ✓ 1846ms | 否 | ✓ 1726ms | ✓ 1783ms | 否 | http |
| 61.52.131.172:8443 | ✓ 1043ms | ✓ 1487ms | ✓ 1167ms | ✓ 1430ms | ✓ 1083ms | http |
| 47.77.216.82:1080 | ✓ 778ms | ✓ 948ms | ✓ 911ms | ✓ 1050ms | ✓ 684ms | http |

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
