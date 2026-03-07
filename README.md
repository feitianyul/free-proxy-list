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

最后更新：2026-03-07 14:01:09 UTC（2026-03-07 22:01:09 UTC+8）

**代理总数：56**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 55 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 56 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 1736ms | 否 | ✓ 1576ms | ✓ 1100ms | ✓ 864ms | http |
| 35.225.22.61:80 | ✓ 824ms | 否 | ✓ 1283ms | ✓ 1099ms | ✓ 1087ms | http |
| 205.209.118.30:3138 | ✓ 269ms | 否 | ✓ 1022ms | ✓ 1070ms | ✓ 1818ms | http |
| 46.249.103.192:443 | ✓ 1638ms | 否 | ✓ 1691ms | ✓ 1633ms | 否 | http |
| 165.227.5.10:8888 | ✓ 1459ms | 否 | ✓ 804ms | ✓ 1178ms | ✓ 1982ms | http |
| 14.143.222.113:10155 | ✓ 1785ms | 否 | ✓ 1043ms | ✓ 1420ms | 否 | http |
| 121.128.121.54:3128 | ✓ 1155ms | 否 | ✓ 1587ms | ✓ 1936ms | 否 | http |
| 168.235.110.63:3128 | ✓ 469ms | 否 | ✓ 753ms | ✓ 1991ms | ✓ 984ms | http |
| 186.148.180.46:999 | ✓ 737ms | ✓ 1624ms | ✓ 809ms | ✓ 1962ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1137ms | 否 | ✓ 1432ms | ✓ 1478ms | ✓ 1181ms | http |
| 14.56.107.244:3128 | ✓ 1817ms | ✓ 1663ms | ✓ 1401ms | 否 | 否 | http |
| 61.72.221.94:3128 | ✓ 1788ms | ✓ 1707ms | 否 | 否 | ✓ 1636ms | http |
| 194.59.204.87:9080 | ✓ 759ms | 否 | ✓ 1173ms | ✓ 1478ms | ✓ 1388ms | http |
| 45.129.141.143:3128 | ✓ 778ms | 否 | ✓ 1862ms | ✓ 1900ms | 否 | http |
| 101.43.255.96:80 | 否 | ✓ 1436ms | ✓ 1588ms | 否 | ✓ 1134ms | http |
| 125.128.12.14:3128 | ✓ 1675ms | ✓ 1771ms | ✓ 1483ms | ✓ 1365ms | ✓ 925ms | http |
| 77.221.18.126:3128 | ✓ 1163ms | ✓ 1715ms | ✓ 1653ms | 否 | ✓ 1782ms | http |
| 61.72.110.94:3128 | ✓ 1690ms | 否 | 否 | ✓ 1603ms | ✓ 1568ms | http |
| 106.14.203.63:3333 | ✓ 1524ms | ✓ 1792ms | 否 | ✓ 1317ms | 否 | http |
| 178.236.245.17:3128 | ✓ 647ms | 否 | ✓ 991ms | 否 | ✓ 1677ms | http |
| 85.9.195.140:1080 | ✓ 769ms | 否 | 否 | ✓ 1390ms | ✓ 1382ms | http |
| 162.248.165.72:1080 | ✓ 1036ms | 否 | ✓ 748ms | 否 | ✓ 1353ms | http |
| 45.136.198.40:3128 | ✓ 668ms | 否 | ✓ 1395ms | ✓ 1992ms | ✓ 1475ms | http |
| 103.215.36.88:18021 | 否 | ✓ 1730ms | ✓ 1695ms | 否 | ✓ 1371ms | http |
| 20.210.76.178:8561 | ✓ 879ms | ✓ 1131ms | ✓ 685ms | ✓ 1025ms | ✓ 764ms | http |
| 20.78.26.206:8561 | ✓ 856ms | ✓ 1041ms | ✓ 821ms | ✓ 1037ms | ✓ 798ms | http |
| 20.27.15.111:8561 | ✓ 872ms | ✓ 1400ms | ✓ 622ms | ✓ 979ms | ✓ 829ms | http |
| 20.27.11.248:8561 | ✓ 874ms | ✓ 1091ms | ✓ 792ms | ✓ 1126ms | ✓ 829ms | http |
| 20.210.39.153:8561 | ✓ 882ms | ✓ 1684ms | ✓ 626ms | ✓ 988ms | ✓ 826ms | http |
| 20.27.14.220:8561 | ✓ 871ms | ✓ 1825ms | ✓ 636ms | ✓ 967ms | ✓ 764ms | http |
| 20.78.118.91:8561 | ✓ 860ms | 否 | ✓ 618ms | ✓ 965ms | ✓ 764ms | http |
| 20.210.76.175:8561 | ✓ 880ms | 否 | ✓ 620ms | ✓ 966ms | ✓ 762ms | http |
| 192.166.82.55:1080 | ✓ 659ms | 否 | ✓ 699ms | ✓ 1621ms | 否 | http |
| 120.92.212.16:7890 | 否 | 否 | ✓ 1147ms | ✓ 1482ms | ✓ 1243ms | http |
| 178.236.245.59:3128 | ✓ 834ms | ✓ 1719ms | ✓ 1981ms | 否 | 否 | http |
| 62.113.119.14:8080 | ✓ 622ms | 否 | 否 | ✓ 1431ms | ✓ 1100ms | http |
| 159.223.42.219:3128 | ✓ 1628ms | 否 | ✓ 976ms | ✓ 1471ms | ✓ 1016ms | http |
| 115.231.181.40:8128 | 否 | 否 | ✓ 1421ms | ✓ 1337ms | ✓ 1055ms | http |
| 138.124.53.25:7443 | ✓ 488ms | 否 | ✓ 1016ms | ✓ 1770ms | ✓ 980ms | http |
| 81.70.169.194:80 | ✓ 1174ms | 否 | ✓ 1165ms | ✓ 1624ms | ✓ 1199ms | http |
| 113.255.59.226:8080 | ✓ 1427ms | 否 | ✓ 1340ms | ✓ 1327ms | ✓ 1349ms | http |
| 207.254.71.62:8088 | ✓ 1045ms | 否 | ✓ 1821ms | ✓ 1563ms | ✓ 1588ms | http |
| 34.101.184.164:3128 | ✓ 1980ms | 否 | 否 | ✓ 1663ms | ✓ 1179ms | http |
| 103.39.51.190:8080 | ✓ 1989ms | 否 | 否 | ✓ 1965ms | ✓ 1803ms | http |
| 46.183.25.8:443 | ✓ 726ms | 否 | ✓ 1173ms | ✓ 1603ms | 否 | http |
| 45.140.147.82:1081 | ✓ 503ms | ✓ 1336ms | ✓ 908ms | ✓ 1119ms | ✓ 854ms | http |
| 45.186.6.104:3128 | ✓ 1485ms | ✓ 1882ms | ✓ 1695ms | 否 | 否 | http |
| 188.132.141.249:443 | ✓ 900ms | 否 | ✓ 1743ms | 否 | ✓ 1612ms | http |
| 14.225.217.30:7890 | ✓ 1031ms | 否 | ✓ 1223ms | 否 | ✓ 1008ms | http |
| 91.193.240.157:9877 | ✓ 1766ms | 否 | ✓ 1522ms | 否 | ✓ 1788ms | http |
| 157.254.37.238:999 | ✓ 1123ms | ✓ 1945ms | ✓ 1658ms | ✓ 1660ms | ✓ 1495ms | http |
| 185.243.218.43:49153 | ✓ 507ms | 否 | ✓ 1420ms | ✓ 1952ms | ✓ 1543ms | http |
| 103.84.95.54:7890 | ✓ 1795ms | 否 | ✓ 1371ms | ✓ 1814ms | ✓ 1673ms | http |
| 88.80.150.82:8080 | ✓ 676ms | 否 | ✓ 1115ms | 否 | ✓ 1241ms | https |
| 180.104.208.71:8081 | ✓ 1131ms | ✓ 1685ms | 否 | ✓ 1424ms | ✓ 1117ms | http |
| 223.16.170.103:80 | ✓ 1326ms | 否 | 否 | ✓ 1297ms | ✓ 1547ms | http |

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
