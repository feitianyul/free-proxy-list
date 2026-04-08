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

最后更新：2026-04-08 14:23:04 UTC（2026-04-08 22:23:04 UTC+8）

**代理总数：64**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 957ms | 否 | ✓ 1126ms | ✓ 1003ms | ✓ 1019ms | http |
| 178.128.24.162:8080 | 否 | 否 | ✓ 1367ms | ✓ 1294ms | ✓ 1041ms | http |
| 159.223.71.162:443 | ✓ 926ms | 否 | 否 | ✓ 1560ms | ✓ 1060ms | http |
| 159.223.71.162:8080 | ✓ 912ms | 否 | 否 | ✓ 1594ms | ✓ 1059ms | http |
| 159.223.213.91:8888 | ✓ 1125ms | 否 | ✓ 1828ms | ✓ 1968ms | ✓ 1988ms | http |
| 167.103.115.102:8800 | ✓ 1388ms | 否 | ✓ 1202ms | 否 | ✓ 1595ms | http |
| 167.103.34.108:8800 | ✓ 1532ms | 否 | ✓ 1617ms | ✓ 1438ms | ✓ 1599ms | http |
| 113.160.132.26:8080 | ✓ 1647ms | 否 | ✓ 1471ms | ✓ 1527ms | ✓ 1408ms | http |
| 111.227.254.12:22222 | ✓ 1286ms | ✓ 1559ms | 否 | 否 | ✓ 1602ms | http |
| 111.227.254.9:22222 | ✓ 1202ms | 否 | ✓ 1227ms | ✓ 1655ms | 否 | http |
| 111.227.254.10:22222 | ✓ 1229ms | 否 | ✓ 1860ms | 否 | ✓ 1262ms | http |
| 111.227.254.11:22222 | ✓ 1244ms | 否 | ✓ 1247ms | ✓ 1550ms | ✓ 1253ms | http |
| 103.252.89.130:8080 | ✓ 649ms | ✓ 1397ms | 否 | ✓ 1696ms | ✓ 1429ms | http |
| 104.234.0.145:55554 | ✓ 824ms | 否 | ✓ 1808ms | ✓ 951ms | ✓ 733ms | http |
| 161.35.70.36:8888 | ✓ 599ms | 否 | ✓ 1242ms | ✓ 1917ms | 否 | http |
| 5.104.87.17:8051 | ✓ 1156ms | 否 | ✓ 1828ms | ✓ 1227ms | 否 | http |
| 167.103.144.127:8800 | ✓ 1786ms | 否 | ✓ 1454ms | 否 | ✓ 1417ms | http |
| 1.231.81.166:3128 | 否 | 否 | ✓ 1664ms | ✓ 1655ms | ✓ 1587ms | http |
| 45.167.124.52:8080 | ✓ 550ms | 否 | ✓ 504ms | ✓ 1532ms | ✓ 1265ms | http |
| 45.167.125.21:999 | ✓ 858ms | 否 | ✓ 1448ms | ✓ 1910ms | ✓ 1457ms | http |
| 46.30.46.133:3128 | ✓ 626ms | ✓ 1440ms | ✓ 536ms | ✓ 1424ms | ✓ 1535ms | http |
| 91.238.104.171:2023 | ✓ 680ms | 否 | ✓ 1500ms | ✓ 1898ms | 否 | http |
| 167.103.31.122:8800 | ✓ 1278ms | 否 | ✓ 1365ms | ✓ 1623ms | ✓ 1431ms | http |
| 59.46.216.131:30001 | ✓ 1132ms | 否 | ✓ 1277ms | 否 | ✓ 1226ms | http |
| 216.24.209.174:4000 | ✓ 75ms | ✓ 930ms | ✓ 281ms | ✓ 995ms | ✓ 723ms | http |
| 216.24.209.3:4000 | ✓ 75ms | 否 | 否 | ✓ 1945ms | ✓ 706ms | http |
| 147.161.239.240:8800 | ✓ 465ms | 否 | ✓ 541ms | ✓ 1472ms | ✓ 1201ms | http |
| 209.38.154.7:1080 | ✓ 1211ms | 否 | 否 | ✓ 1019ms | ✓ 737ms | http |
| 35.225.22.61:80 | ✓ 355ms | 否 | ✓ 614ms | ✓ 1050ms | 否 | http |
| 83.219.250.8:62920 | ✓ 901ms | ✓ 1788ms | ✓ 1454ms | 否 | 否 | http |
| 155.117.18.36:25388 | ✓ 835ms | ✓ 1345ms | ✓ 1350ms | 否 | 否 | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1215ms | ✓ 1614ms | ✓ 1243ms | http |
| 38.92.10.139:33985 | ✓ 1265ms | ✓ 1161ms | ✓ 1853ms | ✓ 954ms | ✓ 894ms | http |
| 38.92.10.152:57579 | ✓ 1259ms | ✓ 1667ms | ✓ 979ms | 否 | 否 | http |
| 43.99.25.221:46509 | ✓ 839ms | ✓ 1373ms | ✓ 868ms | ✓ 1026ms | ✓ 1295ms | http |
| 185.76.241.157:10001 | ✓ 1146ms | 否 | 否 | ✓ 1988ms | ✓ 1514ms | http |
| 154.40.137.209:55965 | ✓ 808ms | ✓ 1948ms | ✓ 1680ms | 否 | 否 | http |
| 218.60.175.168:22222 | 否 | 否 | ✓ 982ms | ✓ 1281ms | ✓ 1031ms | http |
| 91.238.105.64:2024 | 否 | 否 | ✓ 1298ms | ✓ 1954ms | ✓ 1422ms | http |
| 212.58.132.5:8888 | ✓ 1096ms | 否 | ✓ 1794ms | ✓ 1518ms | ✓ 1410ms | http |
| 185.76.240.29:10001 | ✓ 984ms | 否 | ✓ 967ms | 否 | ✓ 1549ms | http |
| 185.76.240.21:10001 | ✓ 961ms | 否 | ✓ 1441ms | ✓ 1811ms | 否 | http |
| 109.107.179.140:8090 | ✓ 734ms | 否 | ✓ 731ms | 否 | ✓ 1517ms | http |
| 194.67.99.223:1080 | ✓ 554ms | 否 | 否 | ✓ 1871ms | ✓ 1625ms | http |
| 34.101.184.164:3128 | ✓ 1729ms | 否 | ✓ 1104ms | ✓ 1697ms | ✓ 1238ms | http |
| 221.122.91.36:11273 | ✓ 983ms | ✓ 1160ms | ✓ 875ms | ✓ 1200ms | 否 | http |
| 221.122.91.36:11195 | ✓ 921ms | ✓ 1165ms | ✓ 928ms | ✓ 1233ms | 否 | http |
| 45.129.141.143:3128 | 否 | 否 | ✓ 1659ms | ✓ 1997ms | ✓ 1676ms | http |
| 209.126.84.232:8888 | ✓ 988ms | 否 | ✓ 1516ms | 否 | ✓ 1781ms | http |
| 20.78.213.56:80 | ✓ 969ms | ✓ 1515ms | ✓ 1770ms | ✓ 1545ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1449ms | ✓ 1448ms | ✓ 1449ms | ✓ 1453ms | ✓ 1145ms | http |
| 47.86.255.82:28314 | ✓ 825ms | ✓ 1150ms | ✓ 1056ms | ✓ 1035ms | ✓ 842ms | http |
| 38.180.2.107:3128 | ✓ 676ms | ✓ 1630ms | ✓ 780ms | 否 | 否 | http |
| 216.167.80.234:58950 | 否 | ✓ 1137ms | ✓ 627ms | ✓ 921ms | ✓ 806ms | http |
| 38.79.118.202:33858 | 否 | 否 | ✓ 1113ms | ✓ 1064ms | ✓ 915ms | http |
| 91.238.104.172:2024 | ✓ 1162ms | ✓ 1524ms | ✓ 1890ms | 否 | ✓ 1757ms | http |
| 91.233.223.147:3128 | ✓ 1190ms | 否 | ✓ 1579ms | 否 | ✓ 1914ms | http |
| 120.92.212.16:8890 | ✓ 1762ms | ✓ 1711ms | ✓ 1428ms | ✓ 1421ms | ✓ 1434ms | http |
| 185.76.240.54:10001 | ✓ 1185ms | 否 | ✓ 937ms | 否 | ✓ 1684ms | http |
| 104.243.46.122:3128 | 否 | ✓ 1228ms | ✓ 1063ms | ✓ 1836ms | ✓ 992ms | http |
| 157.254.37.238:999 | ✓ 1689ms | 否 | 否 | ✓ 1397ms | ✓ 1228ms | http |
| 45.12.151.226:2829 | ✓ 1617ms | 否 | ✓ 695ms | ✓ 1432ms | ✓ 1248ms | http |
| 106.117.208.101:7890 | 否 | ✓ 1509ms | 否 | ✓ 1633ms | ✓ 1385ms | http |
| 146.190.80.158:9090 | ✓ 1346ms | 否 | ✓ 1154ms | ✓ 1548ms | ✓ 1316ms | http |

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
